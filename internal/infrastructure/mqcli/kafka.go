package mqcli

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func newKafka(in digIn) *Kafka {
	return &Kafka{
		config: in.Config.Kafka,
		logger: in.Logger.SysLogger.Named("kafka"),
	}
}

type Kafka struct {
	config *config.KafkaConfig
	logger logger.ILogger
}

func (k *Kafka) GetConfig() *config.KafkaConfig {
	return k.config
}

func (k *Kafka) Produce(ctx context.Context, topic string, value []byte) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(k.config.Host),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(ctx, kafka.Message{
		Value: value,
	})
	if err != nil {
		k.logger.Error(ctx, err)
		return err
	}

	return nil
}

func (k *Kafka) Consume(ctx context.Context, groupID, topic string, handler func(context.Context, []byte) error) {
	if groupID == "" || topic == "" {
		k.logger.Error(ctx, fmt.Errorf("groupID or topic is empty"))
		return
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{k.config.Host},
		GroupID:     groupID,
		GroupTopics: []string{topic},
		MaxBytes:    10e6, // 10MB
	})
	defer reader.Close()

	for {
		select {
		case <-ctx.Done():
			k.logger.Warn(ctx, "server shutdown, kafka listener exit",
				logger.String("topic", topic),
				logger.String("groupID", groupID),
			)
			return
		default:
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				k.logger.Error(ctx, err, zap.Any("message", m))
				k.handleFailedMessage(ctx, topic, m)
			}

			if m.Value == nil {
				continue
			}
			if err := handler(ctx, m.Value); err != nil {
				k.logger.Error(ctx, err, zap.Any("message", m))
				k.handleFailedMessage(ctx, topic, m)
			}
		}
	}
}

func (k *Kafka) handleFailedMessage(ctx context.Context, topic string, msg kafka.Message) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(k.config.Host),
		Topic:    fmt.Sprintf("failed-%s", topic),
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	if err := writer.WriteMessages(ctx, msg); err != nil {
		k.logger.Error(ctx, err, zap.Any("message", msg))
	}
}
