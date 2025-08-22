package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func InfoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	const content = `
	**Katou Megumi (My Istri)**

**Menu: **
!salam
!info
!ask
!editbackground
!jadwalsholat
!sticker
!doa
!jokes
!animequote
!ocr
!asmaulhusna

**Bot created by https://github.com/haikelz/**
`

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
