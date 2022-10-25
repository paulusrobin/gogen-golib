package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	IServer interface {
		Run() error
		Shutdown() error
	}
	kafkaServer struct {
		sig chan os.Signal
		cfg Config
		opt options

		// dependencies
		client sarama.ConsumerGroup
	}
)

// init function to initialize dependencies.
func (s *kafkaServer) init() error {
	saramaClient, err := sarama.NewConsumerGroup(strings.Split(s.cfg.Brokers, ","), s.cfg.Group, s.cfg.SaramaConfig)
	if err != nil {
		return err
	}

	s.client = saramaClient
	return nil
}

// Run function to run http server.
func (s *kafkaServer) Run() error {
	keepRunning := true
	log.Info().Msgf("%s starting a new kafka sarama consumer", s.cfg.logPrefix())

	ctx := context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	consumer := Consumer{ready: make(chan bool), cfg: s.cfg, opt: s.opt}
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalanced happens, the consumer session will need to be
			// recreated to get the new claims
			if err := s.client.Consume(ctx, s.opt.topics, &consumer); err != nil {
				log.Error().Err(err).
					Fields(map[string]interface{}{"options": s.opt}).
					Msgf("%s error kafka sarama consumer", s.cfg.logPrefix())
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Info().Msgf("%s sarama consumer up and running!...", s.cfg.logPrefix())

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Info().Msgf("%s terminating: context cancelled", s.cfg.logPrefix())
			keepRunning = false
		case <-s.sig:
			log.Warn().Msgf("%s terminating: via signal", s.cfg.logPrefix())
			keepRunning = false
		}
	}
	wg.Wait()
	return nil
}

// Shutdown function to close http server.
func (s *kafkaServer) Shutdown() error {
	time.Sleep(s.cfg.GracefulDuration)
	if err := s.client.Close(); err != nil {
		return err
	}
	return nil
}

// Server functions to initialize http server.
func Server(sig chan os.Signal, cfg Config, opts ...Option) (IServer, error) {
	cfg = sanitizeConfig(cfg)
	option := defaultOption
	for _, opt := range opts {
		opt.Apply(&option)
	}

	s := &kafkaServer{sig: sig, cfg: cfg, opt: option}
	if err := s.init(); err != nil {
		log.Error().Err(err).
			Fields(map[string]interface{}{"config": cfg}).
			Msgf("%s failed to initialize server", cfg.logPrefix())
		return nil, err
	}
	return s, nil
}
