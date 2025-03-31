package queue

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"github.com/segmentio/kafka-go"
	"time"
)

type Consumer interface {
	HandleMessage(message string) error
	Listen(ctx context.Context)
	Close() error
}

type consumer struct {
	Reader      *kafka.Reader
	Handler     Consumer
	ServiceName string
}

func (c *consumer) Close() error {
	return c.Reader.Close()
}

func (c *consumer) HandleMessage(message string) error {
	return c.Handler.HandleMessage(message)
}

func (c *consumer) Listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			slg.Logger.Info("Consumer listener is shutting down")
			return
		default:
			ctxRead, cancel := context.WithTimeout(ctx, 5*time.Second)
			msg, err := c.Reader.ReadMessage(ctxRead)
			cancel()
			if err != nil {
				if ctx.Err() != nil {
					slg.Logger.Info("Context cancelled, exiting listen loop")
					return
				}
				slg.Logger.Error("Error reading message", "error", err)
				continue
			}

			slg.Logger.Info("Received message", "message", string(msg.Value))
			if err = c.Handler.HandleMessage(string(msg.Value)); err != nil {
				slg.Logger.Error("Error processing message", "error", err)
			}
		}
	}
}

func NewConsumer(broker, topic, groupId string, handler Consumer) Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &consumer{
		Reader:      reader,
		Handler:     handler,
		ServiceName: "User Service",
	}
}
