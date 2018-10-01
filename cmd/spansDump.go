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
	"encoding/json"
	"expedia.com/haystack-debug-tools/kafkautils"
	pb "expedia.com/haystack-debug-tools/proto"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// spansCmd represents the spans command
var spansCmd = &cobra.Command{
	Use:   "spans",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//set the callback and then run the dump
		kafkaConfig.MsgProcessor = dumpSpans
		kafkautils.RunDumper(kafkaConfig)
	},
}

func init() {
	dumpCmd.AddCommand(spansCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spansCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spansCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func dumpSpans(kafkamsg *kafka.Message, w io.Writer) {
	span := &pb.Span{}
	if err := proto.Unmarshal(kafkamsg.Value, span); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to read protobuf span for partition %d offset %d", kafkamsg.TopicPartition.Partition, kafkamsg.TopicPartition.Offset)
		return
	}
	ajson, err := json.Marshal(span)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to write proto as JSON %s", span.TraceId)
		return
	}
	fmt.Fprintf(w, "%s\n", string(ajson))
}
