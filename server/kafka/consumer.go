package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	cfg   Config
	opt   options
	ready chan bool
}

var defaultHandler = func(cfg Config) HandlerFunc {
	return func(ctx context.Context, message *sarama.ConsumerMessage) error {
		log.Info().
			Fields(map[string]interface{}{"message": message}).
			Msgf("%s running default handler", cfg.logPrefix())
		return nil
	}
}

func (consumer *Consumer) consume(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) {
	var next = defaultHandler(consumer.cfg)
	if consumer.opt.routes != nil {
		if _, exist := consumer.opt.routes[message.Topic]; exist {
			next = consumer.opt.routes[message.Topic]
		}
	}

	if err := next(context.Background(), message); err != nil {
		log.Error().Err(err).Fields(map[string]interface{}{"message": message}).Msgf("error on consuming message")
	} else {
		session.MarkMessage(message, "")
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(group sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			consumer.consume(session, message)

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalancedInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
