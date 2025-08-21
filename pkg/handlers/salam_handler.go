package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func SalamHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, "Assalamu'alaikum")
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
