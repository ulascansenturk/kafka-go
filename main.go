package main

import (
	"fmt"
	"time"

	"github.com/ulascansenturk/go-kafka/kafka"
)

func main() {
	fmt.Println("Kafka starting....")
	go kafka.StartKafka()

	fmt.Println("Kafka Started....")

	time.Sleep(10 * time.Minute)

}
