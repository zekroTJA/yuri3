package config

type Provider interface {
	Load() error
	Instance() *Config
}

type baseProvider struct {
	instance *Config
}

func (p *baseProvider) Instance() *Config {
	return p.instance
}
