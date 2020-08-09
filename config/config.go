package config

import (
	"fmt"

	"github.com/crgimenes/goconfig"
)

type Config struct {
	Amqp AmqpConfig
}

type AmqpConfig struct {
	URL string
}

func LoadConfig() (config Config, err error) {
	goconfig.PrefixEnv = "WAV"

	if err := goconfig.Parse(&config); err != nil {
		return config, err
	}

	fmt.Println(config)

	return config, nil
}
