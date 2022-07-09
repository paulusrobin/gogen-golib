package viper

import (
	consul "github.com/paulusrobin/gogen-golib/consul/interface"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // needed to use viper remote config features
	"time"
)

type consulInstance struct {
	reloader *viper.Viper
	consul.Config
}

func NewConsulReader() {
	viper.New()
}

func (c consulInstance) Read(key string, data interface{}) error {
	param.cfg.setConsulToken()
	if err := c.reloader.AddRemoteProvider(remoteConfigProvider, param.cfg.getConsulRemote(), param.key); err != nil {
		log.Error().Err(err).Msgf("cannot read remote config from consul for key: %s", param.key)
		return err
	}

	c.reloader.SetConfigType(param.configType)
	if err := hotReloadViper.ReadRemoteConfig(); err != nil {
		param.onError(param.key, err)
	} else {
		param.onSuccess(param.key, hotReloadViper)
	}

	log.Info().Msgf("remote config %s was read successfully from consul. starting periodical updates...", param.key)

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			time.Sleep(param.interval)

			if err := hotReloadViper.WatchRemoteConfig(); err != nil {
				param.onError(param.key, err)
				continue
			}

			param.onSuccess(param.key, hotReloadViper)
		}
	}()
	return nil
}
