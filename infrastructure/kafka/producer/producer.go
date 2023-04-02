package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)

	if err != nil {
		fmt.Printf("error", err)
	}
	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		return
	}

	conn.WriteMessages(kafka.Message{Value: []byte("ola gil")})
}
