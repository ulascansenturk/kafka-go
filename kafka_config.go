package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		GroupID:  "t1",
		MaxBytes: 10,
	}
	reader := kafka.NewReader(conf)
	for {
		n, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}

		fmt.Println("Message is :", string(n.Value))

	}
}
