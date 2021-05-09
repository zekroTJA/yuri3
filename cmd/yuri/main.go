package main

import (
	"github.com/joho/godotenv"
	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/yuri3/internal/config"
	"github.com/zekroTJA/yuri3/internal/discord"
	"github.com/zekroTJA/yuri3/internal/static"
)

func main() {
	dotenv(".env", "prod.env", "dev.env")

	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name: static.DiConfigProvider,
		Build: func(ctn di.Container) (interface{}, error) {
			p := config.NewConfitaProvider()
			return p, p.Load()
		},
	})

	builder.Add(di.Def{
		Name: static.DiDiscordProvider,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(static.DiConfigProvider).(config.Provider)
			return discord.NewDiscordgoProvider(cfg.Instance().Discord.Token)
		},
	})

	ctn := builder.Build()
	ctn.Get(static.DiDiscordProvider)
}

func dotenv(optionalFiles ...string) {
	for _, f := range optionalFiles {
		godotenv.Load(f)
	}
}
