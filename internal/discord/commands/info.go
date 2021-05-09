package commands

import "github.com/bwmarrin/discordgo"

type Info struct {
	*discordgo.ApplicationCommand
}

func NewInfo() *Info {
	return &Info{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "info",
			Description: "Basic information",
		},
	}
}

func (c *Info) Exec(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "Heyho!",
		},
	})
}

func (c *Info) Base() *discordgo.ApplicationCommand {
	return c.ApplicationCommand
}
