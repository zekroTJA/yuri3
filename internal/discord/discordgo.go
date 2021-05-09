package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/yuri3/internal/config"
	"github.com/zekroTJA/yuri3/internal/discord/commands"
	"github.com/zekroTJA/yuri3/internal/discord/events"
	"github.com/zekroTJA/yuri3/internal/static"
)

func NewDiscordSession(ctn di.Container) (session *discordgo.Session, err error) {
	cfg := ctn.Get(static.DiConfigProvider).(config.Provider)

	if session, err = discordgo.New("Bot " + cfg.Instance().Discord.Token); err != nil {
		return
	}

	session.AddHandler((&events.Ready{}).Handle)

	cmdHandler := commands.NewHandler(session)
	cmdHandler.Register(commands.NewInfo())

	return
}
