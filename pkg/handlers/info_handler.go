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
	**Katou Megumi**

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
	`

	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
