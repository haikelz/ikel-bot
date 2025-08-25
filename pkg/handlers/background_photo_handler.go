package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func BackgroundPhotoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	// var BACKGROUND_PHOTO_API_URL = os.Getenv("BACKGROUND_PHOTO_API_URL")
	_, err := s.ChannelMessageSend(m.ChannelID, command)
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
