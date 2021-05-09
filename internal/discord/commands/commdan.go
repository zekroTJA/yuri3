package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	Base() *discordgo.ApplicationCommand
	Exec(s *discordgo.Session, i *discordgo.InteractionCreate)
}
