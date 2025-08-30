package utils

import (
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func Get(url string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) []byte {
	response, err := http.Get(url)
	if err != nil {
		MessageWithReply(s, m, "Error fetching data", logger)
		logger.Error("Error fetching data", zap.Error(err))
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		MessageWithReply(s, m, "Error reading data", logger)
		logger.Error("Error reading data", zap.Error(err))
		return nil
	}

	return body
}
