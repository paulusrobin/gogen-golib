package consul

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type (
	// ConnectionConfig object contains consul connection settings.
	ConnectionConfig struct {
		// Host config to set consul host.
		Host string

		// Port config to set consul port.
		Port int

		// Token config to set consul http token.
		Token string

		// Key consul key path to be read.
		Key string
	}

	// Config object to be passed on initialization.
	Config struct {
		// Connection config type to be passed.
		Connection ConnectionConfig

		// ConfigType config type to be passed.
		ConfigType string

		// Interval to run update periodically
		// default 0, it will not run periodically background process to update
		Interval time.Duration
	}

	// Reader interface to read consul settings.
	Reader interface {
		Read(data interface{}) error
	}
)

// SetConsulRemote function to set consul remote host token.
func (c *Config) SetConsulRemote() {
	if os.Getenv("CONSUL_HTTP_TOKEN") == "" && c.Connection.Token != "" {
		_ = os.Setenv("CONSUL_HTTP_TOKEN", c.Connection.Token)
	}
}

// GetConsulRemote function to get consul remote host.
func (c *Config) GetConsulRemote() string {
	if c.Connection.Host == "" {
		c.Connection.Host = os.Getenv("CONSUL_HOST")
	}
	if c.Connection.Port == 0 {
		consulPort := os.Getenv("CONSUL_PORT")
		if consulPort != "" {
			if port, err := strconv.Atoi(consulPort); err == nil {
				c.Connection.Port = port
			}
		}
	}
	return fmt.Sprintf("%s:%d", c.Connection.Host, c.Connection.Port)
}
