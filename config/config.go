package config

import (
	"github.com/crgimenes/goconfig"
)

type Config struct {
	AMQPUrl string
}

func LoadConfig() (config Config, err error) {
	var c Config

	if err := goconfig.Parse(&c); err != nil {
		return config, err
	}

	return c, nil
}
