package consul

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type (
	ConnectionConfig struct {
		Host  string
		Port  int
		Token string
		Key   string
	}
	Config struct {
		Connection ConnectionConfig
		ConfigType string
		Interval   time.Duration
	}
	Reader interface {
		Read(data interface{}) error
	}
)

func (c *Config) SetConsulRemote() {
	if os.Getenv("CONSUL_HTTP_TOKEN") == "" && c.Connection.Token != "" {
		_ = os.Setenv("CONSUL_HTTP_TOKEN", c.Connection.Token)
	}
}

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
