package http

import "fmt"

type HttpConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func (c *HttpConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}