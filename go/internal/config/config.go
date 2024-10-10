package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port              string `envconfig:"PORT" default:":4000"`
	DSN               string `envconfig:"DSN"`
	SessionCookieName string `envconfig:"SESSION_COOKIE_NAME" default:"session"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
