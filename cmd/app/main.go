package main

import (
	"ikel-bot/pkg/configs"
	"ikel-bot/pkg/handlers"
	"ikel-bot/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	logger := configs.NewZap()
	utils.LoadEnv()

	var DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")

	discord := configs.NewDiscord(DISCORD_TOKEN)

	discord.Client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Info("Bot is running")
	})

	discord.Client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == "!info" {
			handlers.InfoHandler(s, m, logger)
		}
		if m.Content == "!ping" {
			handlers.PingHandler(s, m, logger)
		}
		if m.Content == "!salam" {
			handlers.SalamHandler(s, m, logger)
		}
		if m.Content == "!asmaulhusna" {
			handlers.AsmaulHusnaHandler(s, m, logger)
		}
		if m.Content == "!ask" {
			handlers.GeminiHandler(s, m, logger)
		}
		if m.Content == "!jokes" {
			handlers.JokesHandler(s, m, logger)
		}
		if m.Content == "!jadwalsholat" {
			handlers.JadwalSholatHandler(s, m, logger)
		}
	})

	discord.Client.Open()

	defer discord.Client.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
