package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func PingHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			logger.Error("Error sending message", zap.Error(err))
			return
		}
	}
}
