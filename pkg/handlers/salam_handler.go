package handlers

import (
	"ikel-bot/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func SalamHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	utils.MessageWithReply(s, m, "Assalamu'alaikum", logger)
}
