package config

import (
	"github.com/crgimenes/goconfig"
)

type Config struct {
	Amqp AmqpConfig
	DB	 DBConfig
}

type AmqpConfig struct {
	URL string
}

type DBConfig struct {
	User 	 string
	Password string
	Database string
}

func LoadConfig() (config Config, err error) {
	goconfig.PrefixEnv = "WAV"

	if err := goconfig.Parse(&config); err != nil {
		return config, err
	}

	return config, nil
}
