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

func AsmaulHusnaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	var ASMAUL_HUSNA_API_URL = os.Getenv("ASMAUL_HUSNA_API_URL")

	response, err := http.Get(ASMAUL_HUSNA_API_URL + "/api/all")
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

	var asmaulHusnaResponse entities.AsmaulHusnaResponse
	err = json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		logger.Error("Error unmarshalling Asmaul Husna", zap.Error(err))
		return
	}

	loopAsmaulHusnaMessage(0, 20, asmaulHusnaResponse, s, m, logger)
	loopAsmaulHusnaMessage(20, 40, asmaulHusnaResponse, s, m, logger)
	loopAsmaulHusnaMessage(40, 60, asmaulHusnaResponse, s, m, logger)
	loopAsmaulHusnaMessage(60, 80, asmaulHusnaResponse, s, m, logger)
	loopAsmaulHusnaMessage(80, 98, asmaulHusnaResponse, s, m, logger)
}

func loopAsmaulHusnaMessage(start int, end int, asmaulHusnaResponse entities.AsmaulHusnaResponse, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	content := ``

	for _, v := range asmaulHusnaResponse.Data[start:end] {
		content += fmt.Sprintf("%d - %s - %s - %s\n", v.Urutan, v.Latin, v.Arab, v.Arti)
	}

	_, err := s.ChannelMessageSendReply(m.ChannelID, content, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
