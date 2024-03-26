package test_kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

func writeKafka(ctx context.Context) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "image",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("hello world")},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Printf("写入kafka失败: %v\n", err)
			}
		} else {
			break
		}
	}
}

func TestKafka(t *testing.T) {
	c := context.Background()
	writeKafka(c)
}
