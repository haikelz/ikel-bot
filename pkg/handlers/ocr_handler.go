package handlers

import (
	"io"
	"katou-megumi/pkg/utils"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/otiai10/gosseract/v2"
	"go.uber.org/zap"
)

func OcrHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	ocr := gosseract.NewClient()
	defer ocr.Close()
	client := &http.Client{}

	imageUrl := m.Attachments[0].URL

	image, err := client.Get(imageUrl)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting image", logger)
		logger.Error("Error getting image", zap.Error(err))
		return
	}
	defer image.Body.Close()

	imageBytes, err := io.ReadAll(image.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading image", logger)
		logger.Error("Error reading image", zap.Error(err))
		return
	}

	ocr.SetImageFromBytes(imageBytes)

	text, _ := ocr.Text()

	utils.MessageWithReply(s, m, text, logger)
}
