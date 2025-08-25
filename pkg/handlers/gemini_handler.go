package handlers

import (
	"context"
	"ikel-bot/pkg/configs"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"google.golang.org/genai"
)

func GeminiHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var GEMINI_API_KEY = os.Getenv("GEMINI_API_KEY")

	client := configs.NewGemini(context.Background(), GEMINI_API_KEY)

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
		logger.Error("Error generating content", zap.Error(err))
		return
	}

	_, err = s.ChannelMessageSendReply(m.ChannelID, response.Text(), &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
	}
}
