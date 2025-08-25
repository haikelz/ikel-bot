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

	return quoteResponse.Result[0].English
}

func GetQuotes(logger *zap.Logger) string {
	var ANIME_QUOTE_API_URL = os.Getenv("ANIME_QUOTE_API_URL")

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

	return quoteResponse.Result[0].English
}
