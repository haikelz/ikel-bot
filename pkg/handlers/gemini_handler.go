package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func GeminiHandler(s *discordgo.Session, m *discordgo.Message, logger *zap.Logger) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!ask" {
		_, err := s.ChannelMessageSend(m.ChannelID, "")
		if err != nil {
			logger.Error("Error sending message", zap.Error(err))
		}
	}
}
