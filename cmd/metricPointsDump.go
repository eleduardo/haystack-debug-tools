// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/base64"
	"encoding/json"
	"expedia.com/haystack-debug-tools/kafkautils"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"github.com/vmihailenco/msgpack"
	"io"
	"os"
	"strings"
)

type DumpMetricsConfig struct {
	ServiceFilter   string
	OperationFilter string
	EncodingType    string
	Decoder         func(metric string) string
}

var dumpMetricsConfig = &DumpMetricsConfig{}

// metricPointsCmd represents the metricPoints command
var metricPointsCmd = &cobra.Command{
	Use:   "metricPoints",
	Short: "Dumps metric points from the kafka topic as json messages",
	Long: `Metric points in the kafka topic are messagepack encoded so this utility allows for a dump in human
readable format`,
	PreRun: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(dumpMetricsConfig.EncodingType) {
		default:
		case "noop":
			dumpMetricsConfig.Decoder = func(metric string) string {
				return metric
			}
		case "base64":
			dumpMetricsConfig.Decoder = func(metric string) string {
				ab64, err := base64.StdEncoding.WithPadding('_').DecodeString(metric)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Unable to decode string %s", metric)
					return metric
				}
				return string(ab64)
			}
		case "dot":
			dumpMetricsConfig.Decoder = func(metric string) string {
				return strings.Replace(metric, "__", ".", -1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		//set the callback and then run the dump
		kafkaConfig.MsgProcessor = dumpMetricPoints
		kafkautils.RunDumper(kafkaConfig)
	},
}

func init() {
	dumpCmd.AddCommand(metricPointsCmd)
	metricPointsCmd.PersistentFlags().StringVarP(&dumpMetricsConfig.ServiceFilter, "service", "s", "", "Supply this to filter by Service, this operates as a partial case insensitive match")
	metricPointsCmd.PersistentFlags().StringVarP(&dumpMetricsConfig.OperationFilter, "operation", "o", "", "Supply this to filter by Operation within a service, this operates as a partial case insensitive match")
	metricPointsCmd.PersistentFlags().StringVarP(&dumpMetricsConfig.EncodingType, "decoder", "d", "noop", "Encoding of metrics and operation names values are noop, dot, base64")
}

type MetricPoint struct {
	Id        string
	Interval  int8
	Metric    string
	Mtype     string
	Name      string
	OrgId     int8
	Time      uint16
	Value     float64
	Partition int32
	Offset    string
}

func dumpMetricPoints(kafkamsg *kafka.Message, w io.Writer) {
	var item MetricPoint
	msg := kafkamsg.Value
	//unpack the message pack format
	err := msgpack.Unmarshal(msg, &item)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to unmarshal messagepack for partition %s", err)
		return
	}
	//now that we have a struct decorate it with kafka info
	item.Partition = kafkamsg.TopicPartition.Partition
	item.Offset = kafkamsg.TopicPartition.Offset.String()
	item.Metric = runDecoder(item.Metric)
	if item.Metric == "" {
		//this happens because it is being filtered out so don't dump it
		return
	}
	ajson, err := json.Marshal(item)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to unmarshal messagepack for partition %d offset %d", kafkamsg.TopicPartition.Partition, kafkamsg.TopicPartition.Offset)
	}
	fmt.Fprintf(w, "%s\n", string(ajson))
}

func runDecoder(metric string) string {
	tokens := strings.Split(metric, ".")
	var numtokens = len(tokens)
	if numtokens < 4 {
		fmt.Fprintln(os.Stderr, "Unable to unmarshal understand metric string %s", metric)
		return metric
	}
	var sb strings.Builder
	for i := 0; i < len(tokens); i++ {
		if i == 2 {
			//position 2 is the servicename so first clean it
			decoded := dumpMetricsConfig.Decoder(tokens[i])
			if dumpMetricsConfig.ServiceFilter == "" || strings.Contains(strings.ToLower(decoded), strings.ToLower(dumpMetricsConfig.ServiceFilter)) {
				sb.WriteString(decoded)
				sb.WriteString(".")
			} else {
				return ""
			}
		} else if i == 4 {
			//position 4 is the operation name
			decoded := dumpMetricsConfig.Decoder(tokens[i])
			if dumpMetricsConfig.OperationFilter == "" || strings.Contains(strings.ToLower(decoded), strings.ToLower(dumpMetricsConfig.OperationFilter)) {
				sb.WriteString(decoded)
				sb.WriteString(".")
			} else {
				return ""
			}
		} else {
			sb.WriteString(tokens[i])
			sb.WriteString(".")
		}
	}
	return sb.String()
}
