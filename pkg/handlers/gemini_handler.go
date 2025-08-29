package handlers

import (
	"context"
	"errors"
	"katou-megumi/pkg/configs"
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"google.golang.org/genai"
)

func GeminiHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	client := configs.NewGemini(context.Background(), utils.Env().GEMINI_API_KEY)

	response, err := client.Models.GenerateContent(context.Background(), "gemini-2.5-pro", []*genai.Content{
		{
			Role: "user",
			Parts: []*genai.Part{
				{
					Text: "Jawab pertanyaan ini dengan bahasa Indonesia: " + command,
				},
			},
		},
	}, &genai.GenerateContentConfig{
		ResponseMIMEType: "text/plain",
	})
	if err != nil {
		utils.MessageWithReply(s, m, "Error generating content", logger)
		logger.Error("Error generating content", zap.Error(err))
		return
	}

	if response.Text() == "" {
		utils.MessageWithReply(s, m, "Error generating content", logger)
		logger.Error("Error generating content", zap.Error(errors.New("error generating content")))
		return
	}

	utils.MessageWithReply(s, m, response.Text(), logger)
}
