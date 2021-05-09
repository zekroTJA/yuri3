package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/zekroTJA/yuri3/internal/config"
	"github.com/zekroTJA/yuri3/internal/discord"
	"github.com/zekroTJA/yuri3/internal/lavalink"
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
		Name: static.DiDiscord,
		Build: func(ctn di.Container) (interface{}, error) {
			return discord.NewDiscordSession(ctn)
		},
		Close: func(obj interface{}) error {
			logrus.Info("Tearing down Discord connection ...")
			return obj.(*discordgo.Session).Close()
		},
	})

	builder.Add(di.Def{
		Name: static.DiLavalinkProvider,
		Build: func(ctn di.Container) (interface{}, error) {
			return lavalink.NewWaterlinkProvider(ctn), nil
		},
	})

	ctn := builder.Build()
	defer ctn.Delete()

	cfg := ctn.Get(static.DiConfigProvider).(config.Provider)
	logrus.SetLevel(cfg.Instance().Log.Level)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: cfg.Instance().Debug,
	})

	ctn.Get(static.DiLavalinkProvider)
	dc := ctn.Get(static.DiDiscord).(*discordgo.Session)
	if err := dc.Open(); err != nil {
		logrus.WithError(err).Fatal("Failed connecting to discord")
	}
	block()
}

func dotenv(optionalFiles ...string) {
	for _, f := range optionalFiles {
		godotenv.Load(f)
	}
}

func block() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
