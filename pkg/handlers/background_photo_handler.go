package handlers

import (
	"bytes"
	"fmt"
	"ikel-bot/pkg/utils"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func BackgroundPhotoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if len(command) <= 2 {
		if len(command) == 0 {
			logger.Info("User is editing background photo", zap.String("user", m.Author.ID))
			s.ChannelMessageSendReply(m.ChannelID, "Ini adalah perintah untuk mengubah warna background dari sebuah foto. Nama warna yang dimasukkan harus dalam bahasa Inggris.  Contoh: *!editphoto red*", m.Reference())
			return
		}

		utils.MessageWithReply(s, m, "Ini adalah perintah untuk mengubah warna background dari sebuah foto. Nama warna yang dimasukkan harus dalam bahasa Inggris.  Contoh: *!editphoto red*", logger)
		return
	}

	var REMOVE_BG_API_URL = os.Getenv("REMOVE_BG_API_URL")
	var REMOVE_BG_API_KEY = os.Getenv("REMOVE_BG_API_KEY")

	url := fmt.Sprintf("%s/v1.0/removebg", REMOVE_BG_API_URL)

	response, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(m.Embeds[0].Image.URL)))
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}

	response.Header.Set("Accept", "application/json")
	response.Header.Set("Content-Type", "application/json")
	response.Header.Set("X-Api-Key", REMOVE_BG_API_KEY)
}
