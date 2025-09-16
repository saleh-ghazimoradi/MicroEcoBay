package queue

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/domain"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Consume(ctx context.Context, handler func(message domain.Event) error) error
	Close() error
}

type consumer struct {
	reader *kafka.Reader
}

func (c *consumer) Consume(ctx context.Context, handler func(message domain.Event) error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				return fmt.Errorf("failed to read message: %w", err)
			}

			if err := handler(domain.Event{
				Key:   msg.Key,
				Value: msg.Value,
			}); err != nil {
				return fmt.Errorf("failed to handle message: %w", err)
			}

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				return fmt.Errorf("failed to commit message: %w", err)
			}
		}
	}
}

func (c *consumer) Close() error {
	if err := c.reader.Close(); err != nil {
		return fmt.Errorf("failed to close consumer: %w", err)
	}
	return nil
}

func NewConsumer(brokers []string, groupId, topic string) Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &consumer{
		reader: reader,
	}
}
