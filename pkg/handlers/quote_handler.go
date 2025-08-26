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

func QuoteHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if command == "" {
		quotes := GetQuotes(logger)
		_, err := s.ChannelMessageSendReply(m.ChannelID, quotes, &discordgo.MessageReference{
			MessageID: m.ID,
			ChannelID: m.ChannelID,
			GuildID:   m.GuildID,
		})
		if err != nil {
			logger.Error("Error sending message", zap.Error(err))
			return
		}
		return
	}

	quote := GetQuote(command, logger)

	_, err := s.ChannelMessageSendReply(m.ChannelID, quote, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}

func GetQuote(anime string, logger *zap.Logger) string {
	var ANIME_QUOTE_API_URL = os.Getenv("ANIME_QUOTE_API_URL")

	response, err := http.Get(ANIME_QUOTE_API_URL + "/api/getbyanime?anime=" + anime + "&page=1")
	if err != nil {
		logger.Error("Error getting quote", zap.Error(err))
		return ""
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading quote", zap.Error(err))
		return ""
	}

	var quoteResponse entities.QuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		logger.Error("Error unmarshalling quote", zap.Error(err))
		return ""
	}

	content := fmt.Sprintf("%s - %s - %s - %s\n\n", quoteResponse.Result[0].English, quoteResponse.Result[0].Indo, quoteResponse.Result[0].Anime, quoteResponse.Result[0].Character)
	return content
}

func GetQuotes(logger *zap.Logger) string {
	var ANIME_QUOTE_API_URL = os.Getenv("ANIME_QUOTE_API_URL")
	content := ""

	response, err := http.Get(ANIME_QUOTE_API_URL + "/api/getrandom")
	if err != nil {
		logger.Error("Error getting quote", zap.Error(err))
		return ""
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading quote", zap.Error(err))
		return ""
	}

	var quoteResponse entities.QuotesResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		logger.Error("Error unmarshalling quote", zap.Error(err))
		return ""
	}

	for _, v := range quoteResponse.Result {
		content += fmt.Sprintf("%s - %s - %s - %s\n\n", v.English, v.Indo, v.Anime, v.Character)
	}

	return content
}
