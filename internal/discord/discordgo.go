package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/yuri3/internal/config"
	"github.com/zekroTJA/yuri3/internal/discord/commands"
	"github.com/zekroTJA/yuri3/internal/discord/events"
	"github.com/zekroTJA/yuri3/internal/static"
)

type DiscordgoProvider struct {
	session *discordgo.Session
}

func NewDiscordgoProvider(ctn di.Container) (p *DiscordgoProvider, err error) {
	p = &DiscordgoProvider{}

	cfg := ctn.Get(static.DiConfigProvider).(config.Provider)

	if p.session, err = discordgo.New("Bot " + cfg.Instance().Discord.Token); err != nil {
		return
	}

	p.session.AddHandler((&events.Ready{}).Handle)

	cmdHandler := commands.NewHandler(p.session)
	cmdHandler.Register(commands.NewInfo())

	return
}

func (p *DiscordgoProvider) Connect() error {
	return p.session.Open()
}

func (p *DiscordgoProvider) Close() error {
	return p.session.Close()
}
