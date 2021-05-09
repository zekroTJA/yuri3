package commands

import "github.com/bwmarrin/discordgo"

type Handler struct {
	session *discordgo.Session
	cmds    map[string]Command
}

func NewHandler(session *discordgo.Session) (h *Handler) {
	h = &Handler{
		session: session,
		cmds:    make(map[string]Command),
	}

	session.AddHandler(func(s *discordgo.Session, e *discordgo.InteractionCreate) {
		if c, ok := h.cmds[e.Data.Name]; ok {
			c.Exec(s, e)
		}
	})

	session.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		for _, c := range h.cmds {
			h.session.ApplicationCommandCreate(h.session.State.User.ID, "", c.Base())
		}
	})

	return
}

func (h *Handler) Register(c Command) {
	h.cmds[c.Base().Name] = c
}
