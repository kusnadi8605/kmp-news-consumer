package config

import (
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

//KafkaConn configuration
func KafkaConn(broker []string, topic string, group string) *kafka.Reader {

	kafconf := kafka.ReaderConfig{
		Brokers:         broker,
		GroupID:         group,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
		CommitInterval:  time.Second,
	}

	// to consume messages
	fmt.Printf("kafka start consuming topic : %v", topic)
	//log.Logf("kafka start consuming topic : %v", topic)

	reader := kafka.NewReader(kafconf)

	return reader
}
