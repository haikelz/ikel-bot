package handlers

import (
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func OcrHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	// ocr := configs.NewGoserract()

	utils.MessageWithReply(s, m, "OCR", logger)
}
