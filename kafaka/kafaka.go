package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	getMessage()

}
func sendMessage() {
	// to produce messages
	topic := "my-topic"
	partition := 0
	conn, error := kafka.DialLeader(context.Background(), "tcp", "192.168.99.100:9092", topic, partition)

	if error != nil {
		log.Fatalln(error)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	i, error := conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if error != nil {
		log.Fatalln(error)
	}
	log.Println(i)

	conn.Close()

}
func getMessage() {
	// to consume messages
	topic := "my-topic"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "192.168.99.100:9093", topic, partition)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		log.Println(string(b))
	}

	batch.Close()
	conn.Close()

}
