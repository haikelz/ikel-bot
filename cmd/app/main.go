package main

import (
	"ikel-bot/pkg/handlers"
	"ikel-bot/pkg/services"
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

	discordService := services.NewDiscordService(token)

	discordService.Client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Info("Bot is running")
	})

	discordService.Client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.PingHandler(s, m, logger)
	})
	discordService.Client.Open()

	defer discordService.Client.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
