package discord

import "github.com/bwmarrin/discordgo"

type DiscordgoProvider struct {
	session *discordgo.Session
}

func NewDiscordgoProvider(token string) (p *DiscordgoProvider, err error) {
	p = &DiscordgoProvider{}

	p.session, err = discordgo.New("bot " + token)

	return
}
