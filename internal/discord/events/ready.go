package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Ready struct{}

func (r *Ready) Handle(s *discordgo.Session, e *discordgo.Ready) {
	logrus.WithField("id", s.State.User.ID).Info("Connection to Discord is ready")
}
