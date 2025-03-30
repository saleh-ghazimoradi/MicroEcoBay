package queue

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, key, value []byte) error
	Close() error
}

type producer struct {
	writer *kafka.Writer
}

func (p *producer) PublishMessage(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (p *producer) Close() error {
	return p.writer.Close()
}

func createTopic(broker, topic string) error {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return err
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return err
	}

	for _, p := range partitions {
		slg.Logger.Info("Partition found", "partition", p)
		if p.Topic == topic {
			return nil
		}
	}

	return conn.CreateTopics(
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
}

func NewProducer(broker, topic string) Producer {
	if err := createTopic(broker, topic); err != nil {
		slg.Logger.Error("Error creating topic", "error", err)
	}
	return &producer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(broker),
			Topic: topic,
		},
	}
}
