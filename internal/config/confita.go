package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

type ConfitaProvider struct {
	*baseProvider
}

func NewConfitaProvider() *ConfitaProvider {
	return &ConfitaProvider{
		baseProvider: newBaseProvider(),
	}
}

func (p *ConfitaProvider) Load() error {
	return confita.NewLoader(
		file.NewOptionalBackend("./config.json"),
		file.NewOptionalBackend("./config.yml"),
		env.NewBackend(),
		flags.NewBackend(),
	).Load(context.Background(), p.instance)
}
