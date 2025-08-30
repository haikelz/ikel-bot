package handlers

import (
	"encoding/json"
	"fmt"
	"katou-megumi/pkg/entities"
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func DoaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	body := utils.Get(utils.Env().DOA_API_URL+"/api/doa/v1/random", s, m, logger)

	var doaResponse []entities.Doa
	err := json.Unmarshal(body, &doaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling do'a", logger)
		logger.Error("Error unmarshalling do'a", zap.Error(err))
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", doaResponse[0].ID, doaResponse[0].Doa, doaResponse[0].Ayat, doaResponse[0].Artinya, command), logger)
}
