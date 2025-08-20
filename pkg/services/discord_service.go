package services

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordService struct {
	Client *discordgo.Session
}

func NewDiscordService(token string) *DiscordService {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Failed to create Discord client: %v", err)
		return nil
	}

	return &DiscordService{
		Client: client,
	}
}

func (s *DiscordService) SendMessage(channelID string, message string) error {
	_, err := s.Client.ChannelMessageSend(channelID, message)
	return err
}
