package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func PingHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	_, err := s.ChannelMessageSendReply(m.ChannelID, "Pong!", &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
