// Пакет для работы с очередью сообщений.
package queue

import (
	"errors"

	kafka "github.com/segmentio/kafka-go"
)

// Client - клиент очереди Kafka.
type Client struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

// New создает и инициализирует клиента Kafka.
func New(brokers []string, topic string, groupId string) (*Client, error) {
	if len(brokers) < 1 {
		return nil, errors.New("не указаны брокеры")
	}
	c := Client{}
	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})
	c.Writer = &kafka.Writer{
		Addr:     kafka.TCP(brokers[0]),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &c, nil
}
