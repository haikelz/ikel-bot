package handlers

import (
	"encoding/json"
	"ikel-bot/pkg/entities"
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func AsmaulHusnaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	var ASMAUL_HUSNA_API_URL = os.Getenv("ASMAUL_HUSNA_API_URL")

	response, err := http.Get(ASMAUL_HUSNA_API_URL)
	if err != nil {
		logger.Error("Error fetching Asmaul Husna", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading Asma'ul Husna", zap.Error(err))
		return
	}

	var asmaulHusna []entities.AsmaulHusna
	err = json.Unmarshal(body, &asmaulHusna)
	if err != nil {
		logger.Error("Error unmarshalling Asmaul Husna", zap.Error(err))
		return
	}

	const content = `
	
	`

	_, err = s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}

	/*response, err := http.Get(ASMAUL_HUSNA_API_URL)
	if err != nil {
		logger.Error("Error fetching Asmaul Husna", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading Asmaul Husna", zap.Error(err))
		return
	}

	var asmaulHusna []map[string]interface{}
	err = json.Unmarshal(body, &asmaulHusna)
	if err != nil {
		logger.Error("Error unmarshalling Asmaul Husna", zap.Error(err))
		return
	}

	randomIndex := rand.Intn(len(asmaulHusna))
	randomAsmaulHusna := asmaulHusna[randomIndex]

	logger.Info("Asmaul Husna", zap.Any("asmaulHusna", randomAsmaulHusna))
	logger.Info("Asmaul Husna", zap.Any("asmaulHusna", randomAsmaulHusna["name"]))*/
}
