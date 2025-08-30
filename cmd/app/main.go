package main

import (
	"katou-megumi/pkg/configs"
	"katou-megumi/pkg/handlers"
	"katou-megumi/pkg/utils"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	logger := configs.NewZap()

	discord := configs.NewDiscord(utils.Env().DISCORD_TOKEN)

	discord.Client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		handlers.UserInfoHandler(s, r, logger)
	})

	discord.Client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		var splitMessage []string = strings.Split(m.Content, " ")
		var args []string = splitMessage[1:]
		var command string = strings.Join(args, " ")

		if splitMessage[0] == "!info" {
			handlers.InfoHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!ping" {
			handlers.PingHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!salam" {
			handlers.SalamHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!asmaulhusna" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.AsmaulHusnaHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!ask" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.GeminiHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!jokes" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.JokeHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!jadwalsholat" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.JadwalSholatHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!doa" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.DoaHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!quote" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.QuoteHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!editbackground" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.BackgroundPhotoHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!ocr" {
			s.ChannelMessageSend(m.ChannelID, utils.WAIT_MESSAGE)
			handlers.OcrHandler(s, m, logger, command)
		}
	})

	discord.Client.Open()

	defer discord.Client.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
