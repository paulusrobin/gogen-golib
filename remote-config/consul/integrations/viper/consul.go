package viper

import (
	consul "github.com/paulusrobin/gogen-golib/remote-config/consul/interface"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // needed to use viper remote config features
	"time"
)

const configProvider = `consul`

type consulInstance struct {
	cfg      consul.Config
	reloader *viper.Viper
}

// NewConsulReader functions to initialize consul.Reader.
func NewConsulReader(cfg consul.Config) (consul.Reader, error) {
	reloader := viper.New()
	cfg.SetConsulRemote()
	if err := reloader.AddRemoteProvider(configProvider, cfg.GetConsulRemote(), cfg.Connection.Key); err != nil {
		log.Error().Err(err).Msgf("cannot read remote config from consul for key: %s", cfg.Connection.Key)
		return nil, err
	}
	reloader.SetConfigType(cfg.ConfigType)
	return &consulInstance{
		cfg:      cfg,
		reloader: reloader,
	}, nil
}

// Read functions to read consul data implement consul.Reader.
func (c consulInstance) Read(data interface{}) error {
	if err := c.reloader.ReadRemoteConfig(); err != nil {
		log.Error().Err(err).
			Fields(map[string]interface{}{"config": c.cfg}).
			Msgf("cannot read remote config")
		return err
	}

	if err := c.reloader.Unmarshal(&data); err != nil {
		log.Error().Err(err).
			Fields(map[string]interface{}{"config": c.cfg}).
			Msgf("cannot unmarshall remote config to data")
		return err
	}

	log.Info().
		Fields(map[string]interface{}{"config": c.cfg, "data": data}).
		Msgf("success read remote config %s from consul", c.cfg.Connection.Key)

	if c.cfg.Interval <= 0 {
		return nil
	}

	log.Info().Msgf("starting remote config %s periodical updates...", c.cfg.Connection.Key)

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			time.Sleep(c.cfg.Interval)

			if err := c.reloader.WatchRemoteConfig(); err != nil {
				log.Warn().Err(err).
					Fields(map[string]interface{}{"config": c.cfg, "data": data}).
					Msgf("cannot watch remote config, using previous config")
				continue
			}

			if err := c.reloader.Unmarshal(&data); err != nil {
				log.Warn().Err(err).
					Fields(map[string]interface{}{"config": c.cfg, "data": data}).
					Msgf("cannot unmarshall periodical remote config to data, using previous config")
				continue
			}
		}
	}()
	return nil
}
