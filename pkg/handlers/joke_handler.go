package handlers

import (
	"encoding/json"
	"io"
	"katou-megumi/pkg/entities"
	"katou-megumi/pkg/utils"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func JokeHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	var JOKES_API_URL = os.Getenv("JOKES_API_URL")

	imageUrl, err := getJokeImage(JOKES_API_URL, s, m, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting joke image", logger)
		logger.Error("Error getting joke image", zap.Error(err))
		return
	}

	jokeText, err := getJokeText(JOKES_API_URL, s, m, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting joke text", logger)
		logger.Error("Error getting joke text", zap.Error(err))
		return
	}

	utils.MessageWithEmbedReply(s, m, &discordgo.MessageEmbed{
		Title: jokeText,
		Image: &discordgo.MessageEmbedImage{URL: imageUrl},
	}, logger)
}

func getJokeImage(JOKES_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) (string, error) {
	response, err := http.Get(JOKES_API_URL + "/api/image/random")
	if err != nil {
		utils.MessageWithReply(s, m, "Error fetching jokes", logger)
		logger.Error("Error fetching jokes", zap.Error(err))
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading response body", logger)
		logger.Error("Error reading response body", zap.Error(err))
		return "", err
	}

	var jokeImageResponse entities.JokeImageResponse
	err = json.Unmarshal(body, &jokeImageResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling jokes", logger)
		logger.Error("Error unmarshalling jokes", zap.Error(err))
	}

	return jokeImageResponse.Data.Url, nil
}

func getJokeText(JOKES_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) (string, error) {
	response, err := http.Get(JOKES_API_URL + "/api/text/random")
	if err != nil {
		utils.MessageWithReply(s, m, "Error fetching jokes", logger)
		logger.Error("Error fetching jokes", zap.Error(err))
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading response body", zap.Error(err))
		return "", err
	}

	var jokeTextResponse entities.JokeTextResponse
	err = json.Unmarshal(body, &jokeTextResponse)
	if err != nil {
		logger.Error("Error unmarshalling jokes", zap.Error(err))
	}

	return jokeTextResponse.Data, nil
}
