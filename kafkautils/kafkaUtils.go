package kafkautils

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"io"
	"os"
	"os/signal"
	"syscall"
)

type KafkaDumper struct {
	Topic        string
	BootStrap    string
	MsgProcessor func(msg *kafka.Message, w io.Writer)
}

func RunDumper(dumper *KafkaDumper) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	group := "haystack-debug" + uuid.New().String()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               dumper.BootStrap,
		"group.id":                        group,
		"enable.auto.commit":              true,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer:", err)
		os.Exit(1)
	}
	fmt.Printf("Created Consumer %v with group\n", consumer)
	err = consumer.Subscribe(dumper.Topic, nil)
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		case ev := <-consumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				consumer.Assign(e.Partitions)

			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				consumer.Unassign()
			case *kafka.Message:
				dumper.MsgProcessor(e, os.Stdout)
			case kafka.PartitionEOF:
				fmt.Fprintf(os.Stderr, "%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			}
		}
	}
	fmt.Printf("Closing consumer\n")
	consumer.Close()
}
