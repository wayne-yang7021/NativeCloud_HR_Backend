package messagequeue

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/4040www/NativeCloud_HR/kafka/model"
	"github.com/Shopify/sarama"
)

func StartConsumer(brokers string, groupID string) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_1_0_0

	consumerGroup, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), groupID, config)
	if err != nil {
		return err
	}

	ctx := context.Background()

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{CheckInTopic}, &consumerGroupHandler{}); err != nil {
				log.Printf("Error from consumer: %v", err)
			}
		}
	}()

	return nil
}

type consumerGroupHandler struct{}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var record model.CheckInRequest
		if err := json.Unmarshal(message.Value, &record); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		if err := model.CreateCheckInRecord(&record); err != nil {
			log.Printf("Failed to save to DB: %v", err)
		}

		session.MarkMessage(message, "")
	}
	return nil
}
