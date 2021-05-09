package lavalink

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lukasl-dev/waterlink"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/zekroTJA/yuri3/internal/config"
	"github.com/zekroTJA/yuri3/internal/static"
)

type WaterlinkProvider struct {
	client *waterlink.Client
}

func NewWaterlinkProvider(ctn di.Container) (p *WaterlinkProvider) {
	p = &WaterlinkProvider{}

	session := ctn.Get(static.DiDiscord).(*discordgo.Session)
	cfg := ctn.Get(static.DiConfigProvider).(config.Provider)

	session.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		var err error
		p.client, err = waterlink.New(
			waterlink.HTTP(cfg.Instance().Lavalink.HttpAddress),
			waterlink.WS(cfg.Instance().Lavalink.WSAddress),
			waterlink.Password(cfg.Instance().Lavalink.Password),
			waterlink.UserID(s.State.User.ID),
		)
		if err != nil {
			logrus.WithError(err).Fatal("Failed initializing lavalink connection")
		} else {
			logrus.Info("Lavalink connection established")
		}
	})

	session.AddHandler(func(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
		err := p.client.VoiceUpdate(e.GuildID, s.State.SessionID, waterlink.VoiceServerUpdate{
			GuildID:  e.GuildID,
			Token:    e.Token,
			Endpoint: e.Endpoint,
		})
		if err != nil {
			logrus.WithError(err).WithField("gid", e.GuildID).Error("Failed updating voice state")
		}
	})

	return
}
