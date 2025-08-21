package main

import (
	"ikel-bot/pkg/configs"
	"ikel-bot/pkg/handlers"
	"ikel-bot/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	utils.LoadEnv()

	token := os.Getenv("DISCORD_TOKEN")

	discord := configs.NewDiscord(token)

	discord.Client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Info("Bot is running")
	})

	discord.Client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.PingHandler(s, m, logger)
	})
	discord.Client.Open()

	defer discord.Client.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
