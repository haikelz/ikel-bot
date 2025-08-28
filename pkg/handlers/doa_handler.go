package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"katou-megumi/pkg/entities"
	"katou-megumi/pkg/utils"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func DoaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	var DOA_API_URL = os.Getenv("DOA_API_URL")

	response, err := http.Get(DOA_API_URL + "/api/doa/v1/random")
	if err != nil {
		utils.MessageWithReply(s, m, "Error fetching do'a", logger)
		logger.Error("Error fetching do'a", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading do'a", logger)
		logger.Error("Error reading do'a", zap.Error(err))
		return
	}

	var doaResponse []entities.Doa
	err = json.Unmarshal(body, &doaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling do'a", logger)
		logger.Error("Error unmarshalling do'a", zap.Error(err))
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", doaResponse[0].ID, doaResponse[0].Doa, doaResponse[0].Ayat, doaResponse[0].Artinya, command), logger)
}
