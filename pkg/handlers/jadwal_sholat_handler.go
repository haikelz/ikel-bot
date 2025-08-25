package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func JadwalSholatHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	//var QURAN_API_URL = os.Getenv("QURAN_API_URL")
	_, err := s.ChannelMessageSendReply(m.ChannelID, command, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
