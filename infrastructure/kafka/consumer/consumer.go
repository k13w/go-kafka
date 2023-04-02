package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)

	fmt.Println("consumer", consumer)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing consumer: %v", err)
		}
	}()

	fmt.Println("buuh")

	partitionConsumer, err := consumer.ConsumePartition("quickstart-events", 0, sarama.OffsetNewest)

	fmt.Println("oii consumer")
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}

	fmt.Println("dsadsdsa")
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	fmt.Println("batch")

	batchSize := 100
	messages := make([]*sarama.ConsumerMessage, 0, batchSize)

	fmt.Println("loop", partitionConsumer.Messages())

	for message := range partitionConsumer.Messages() {
		fmt.Println("aoba")
		messages = append(messages, message)
		fmt.Println("men", message)
		fmt.Println("mens", messages)
		if len(messages) == batchSize {
			processBatch(messages)
			messages = make([]*sarama.ConsumerMessage, 0, batchSize)

			fmt.Println("mensagens", messages)
		}
	}
}

func processBatch(messages []*sarama.ConsumerMessage) {
	for _, message := range messages {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}
}
