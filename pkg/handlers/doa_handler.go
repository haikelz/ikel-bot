package handlers

import (
	"encoding/json"
	"fmt"
	"ikel-bot/pkg/entities"
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func DoaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	var DOA_API_URL = os.Getenv("DOA_API_URL")

	response, err := http.Get(DOA_API_URL)
	if err != nil {
		logger.Error("Error fetching do'a", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading do'a", zap.Error(err))
		return

	}

	var doaResponse []entities.Doa
	err = json.Unmarshal(body, &doaResponse)
	if err != nil {
		logger.Error("Error unmarshalling do'a", zap.Error(err))
		return
	}

	loopDoaMessage(0, 10, doaResponse, s, m, logger)
}

func loopDoaMessage(start int, end int, doaResponse []entities.Doa, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	content := ``

	for _, v := range doaResponse[start:end] {
		content += fmt.Sprintf("%d - %s - %s - %s\n", v.ID, v.Doa, v.Ayat, v.Artinya)
	}

	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
