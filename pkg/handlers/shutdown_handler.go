package handlers

import (
	"katou-megumi/pkg/utils"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func ShutdownHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	utils.MessageWithReply(s, m, "Shutting down...", logger)

	go func() {
		time.Sleep(1 * time.Second)

		s.Close()
		os.Exit(0)
	}()
}
