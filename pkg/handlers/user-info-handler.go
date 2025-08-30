package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func UserInfoHandler(s *discordgo.Session, r *discordgo.Ready, logger *zap.Logger) {
	log.Printf("Logged in as: %v#%v", r.User.Username, r.User.Discriminator)
}
