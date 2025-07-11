package queue

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/slg"
	"github.com/segmentio/kafka-go"
	"time"
)

type Consumer struct {
	catalogService service.CatalogService
	reader         *kafka.Reader
}

func (c *Consumer) HandleMessage(message string) error {
	slg.Logger.Info("Received kafka messages inside product service: %s", message)
	return nil
}

func (c *Consumer) Listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			slg.Logger.Info("Consumer listener is shutting down")
			return
		default:
			ctxRead, cancel := context.WithTimeout(ctx, 5*time.Second)
			msg, err := c.reader.ReadMessage(ctxRead)
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
			if err = c.HandleMessage(string(msg.Value)); err != nil {
				slg.Logger.Error("Error processing message", "error", err)
			}
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func NewConsumer(catalogService service.CatalogService, broker, topic, groupId string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &Consumer{
		catalogService: catalogService,
		reader:         reader,
	}
}
