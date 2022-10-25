package kafka

import (
	"github.com/Shopify/sarama"
	"time"
)

const (
	defaultGracefulDuration = 30 * time.Second
)

// Config http config server.
type Config struct {
	// Brokers required, kafka brokers separated by comma.
	Brokers string

	// Group required, kafka consumer group.
	Group string

	// SaramaConfig required, kafka sarama config.
	SaramaConfig *sarama.Config

	// Name optional, kafka server name to be exposed.
	Name string

	// GracefulDuration optional, graceful duration to shut down the server, 30 seconds by default.
	GracefulDuration time.Duration
}

func (c Config) logPrefix() string {
	if c.Name != "" {
		return `[kafka-server: ` + c.Name + `]`
	}
	return `[kafka-server]`
}

func sanitizeConfig(cfg Config) Config {
	if cfg.GracefulDuration == 0 {
		cfg.GracefulDuration = defaultGracefulDuration
	}
	return cfg
}
