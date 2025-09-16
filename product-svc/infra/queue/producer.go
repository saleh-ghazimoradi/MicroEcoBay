package queue

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"github.com/segmentio/kafka-go"
	"time"
)

type Producer interface {
	Produce(ctx context.Context, event domain.Event) error
	Close() error
}

type producer struct {
	writer *kafka.Writer
}

func (p *producer) Produce(ctx context.Context, event domain.Event) error {
	if err := p.writer.WriteMessages(ctx, kafka.Message{
		Key:   event.Key,
		Value: event.Value,
	}); err != nil {
		return fmt.Errorf("failed to produce event: %w", err)
	}
	return nil
}

func (p *producer) Close() error {
	if err := p.writer.Close(); err != nil {
		return fmt.Errorf("failed to close producer: %w", err)
	}
	return nil
}

func NewProducer(brokers []string, topic string) Producer {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		RequiredAcks:           kafka.RequireAll,
		MaxAttempts:            3,
		WriteTimeout:           10 * time.Second,
		ReadTimeout:            10 * time.Second,
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
	}
	return &producer{
		writer: writer,
	}
}
