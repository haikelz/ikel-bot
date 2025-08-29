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

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", r.User.Username, r.User.Discriminator)
	})

	return &Discord{
		Client: client,
	}
}
