package consul

import (
	"fmt"
	"os"
	"strconv"
)

type (
	Config struct {
		Host  string
		Port  int
		Token string
	}
	Reader interface {
		Read(key string, data interface{}) error
	}
)

func (c *Config) GetConsulRemote() string {
	if c.Host == "" {
		c.Host = os.Getenv("CONSUL_HOST")
	}
	if c.Port == 0 {
		consulPort := os.Getenv("CONSUL_PORT")
		if consulPort != "" {
			if port, err := strconv.Atoi(consulPort); err == nil {
				c.Port = port
			}
		}
	}
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
