package config

import "github.com/sirupsen/logrus"

var defaultConfig = Config{
	Debug: false,
	Log: Log{
		Level: logrus.InfoLevel,
	},
	Discord: Discord{},
}

type Provider interface {
	Load() error
	Instance() *Config
}

type baseProvider struct {
	instance *Config
}

func newBaseProvider() *baseProvider {
	return &baseProvider{
		instance: &defaultConfig,
	}
}

func (p *baseProvider) Instance() *Config {
	return p.instance
}
