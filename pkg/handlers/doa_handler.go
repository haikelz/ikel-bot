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

func DoaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	var DOA_API_URL = os.Getenv("DOA_API_URL")

	response, err := http.Get(DOA_API_URL + "/api/doa/v1/random")
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

	_, err = s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("%d - %s - %s - %s", doaResponse[0].ID, doaResponse[0].Doa, doaResponse[0].Ayat, doaResponse[0].Artinya, command), &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
