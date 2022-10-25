package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

type (
	HandlerFunc func(ctx context.Context, message *sarama.ConsumerMessage) error

	Option interface {
		Apply(o *options)
	}
	options struct {
		topics []string
		routes map[string]HandlerFunc
	}
)

var defaultOption = options{
	topics: make([]string, 0),
	routes: make(map[string]HandlerFunc),
}
