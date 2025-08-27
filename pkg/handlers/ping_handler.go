package handlers

import (
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func PingHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	utils.MessageWithReply(s, m, "Pong!", logger)
}
