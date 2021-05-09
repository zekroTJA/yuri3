package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/yuri3/internal/discord/events"
)

type DiscordgoProvider struct {
	session *discordgo.Session
}

func NewDiscordgoProvider(token string) (p *DiscordgoProvider, err error) {
	p = &DiscordgoProvider{}

	p.session, err = discordgo.New("Bot " + token)

	p.session.AddHandler((&events.Ready{}).Handle)

	return
}

func (p *DiscordgoProvider) Connect() error {
	return p.session.Open()
}

func (p *DiscordgoProvider) Close() error {
	return p.session.Close()
}
