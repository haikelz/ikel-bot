package utils

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// An utility function for discord's sending message with zap logger

// Reply user's message
func MessageWithReply(s *discordgo.Session, m *discordgo.MessageCreate, content string, logger *zap.Logger) {
	_, err := s.ChannelMessageSendReply(m.ChannelID, content, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}

// Send message without mentioned the user
func Message(s *discordgo.Session, m *discordgo.MessageCreate, content string, logger *zap.Logger) {
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}

// Send message with embed reply. Suitable for sending image or video.
func MessageWithEmbedReply(s *discordgo.Session, m *discordgo.MessageCreate, embed *discordgo.MessageEmbed, logger *zap.Logger) {
	_, err := s.ChannelMessageSendEmbedReply(m.ChannelID, embed, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
