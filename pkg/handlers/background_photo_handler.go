package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"katou-megumi/pkg/entities"
	"katou-megumi/pkg/utils"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func BackgroundPhotoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if command == "info" {
		utils.MessageWithReply(s, m, "Ini adalah perintah untuk mengubah warna background dari sebuah foto. Nama warna yang dimasukkan harus dalam bahasa Inggris.  Contoh: *!editphoto red*", logger)
		return
	}

	imageUrl := m.Attachments[0].URL

	removeBgData := entities.RemoveBgRequest{ImageUrl: imageUrl, BgColor: command}

	jsonData, err := json.Marshal(removeBgData)
	if err != nil {
		utils.MessageWithReply(s, m, "Error marshalling data", logger)
		logger.Error("Error marshalling data", zap.Error(err))
		return
	}

	url := fmt.Sprintf("%s/v1.0/removebg", utils.Env().REMOVE_BG_API_URL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))

	if err != nil {
		utils.MessageWithReply(s, m, "Error sending message", logger)
		logger.Error("Error sending message", zap.Error(err))
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", utils.Env().REMOVE_BG_API_KEY)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		utils.MessageWithReply(s, m, "Error sending message", logger)
		logger.Error("Error sending message", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading response body", logger)
		logger.Error("Error reading response body", zap.Error(err))
		return
	}

	utils.MessageWithEmbedReply(s, m, &discordgo.MessageEmbed{
		Title: "Background Photo",
		Image: &discordgo.MessageEmbedImage{URL: string(body)},
	}, logger)
}
