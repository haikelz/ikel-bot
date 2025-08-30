package configs

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	Client *discordgo.Session
}

func NewDiscord(token string) *Discord {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Failed to create Discord client: %v", err)
		return nil
	}

	client.StateEnabled = true

	return &Discord{
		Client: client,
	}
}
