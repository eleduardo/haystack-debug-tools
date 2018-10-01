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
	"expedia.com/haystack-debug-tools/kafkautils"
	"fmt"

	"github.com/spf13/cobra"
	"os"
)

var kafkaConfig = &kafkautils.KafkaDumper{}

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "kafkaDump",
	Short: "Dump content from a Haystack kafka topic",
	Long: `With this command you can dump data from different kafka topics, it will decode the particular stream into
a human readable set of data.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if kafkaConfig.BootStrap == "" {
			fmt.Fprintf(os.Stderr, "A bootstrap connection string is required to connect to Kakfa")
			os.Exit(1)
		}
		if kafkaConfig.Topic == "" {
			fmt.Fprintf(os.Stderr, "A topic to listen to is required")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	rootCmd.PersistentFlags().StringVarP(&kafkaConfig.BootStrap, "bootstrap", "b", "", "Kafka Bootstrap connection string")
	rootCmd.PersistentFlags().StringVarP(&kafkaConfig.Topic, "topic", "t", "", "Topic to consumer from")
}
