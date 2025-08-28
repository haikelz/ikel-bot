package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"katou-megumi/pkg/entities"
	"katou-megumi/pkg/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func AsmaulHusnaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	var ASMAUL_HUSNA_API_URL = os.Getenv("ASMAUL_HUSNA_API_URL")
	/*
		TODO:
		- Buat case jika input dari user itu urutan asmaul husna
		- Buat case jika input dari user itu nama asmaul husna
	*/

	if command == "" {
		asmaulHusnaResponse := getAllAsmaulHusna(ASMAUL_HUSNA_API_URL, s, m, logger)

		if asmaulHusnaResponse == nil {
			logger.Error("Error fetching Asmaul Husna", zap.Error(errors.New("error fetching Asmaul Husna")))
			return
		}

		loopAsmaulHusnaMessage(0, 20, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(20, 40, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(40, 60, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(60, 80, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(80, 98, asmaulHusnaResponse, s, m, logger)

		return
	}

	if number, err := strconv.Atoi(command); err == nil {
		asmaulHusnaResponse := getAsmaulHusnaByUrutan(number, ASMAUL_HUSNA_API_URL, s, m, logger)

		if asmaulHusnaResponse.Data.Urutan == 0 {
			utils.MessageWithReply(s, m, "Asmaul Husna tidak ditemukan", logger)
			return
		}

		utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", asmaulHusnaResponse.Data.Urutan, asmaulHusnaResponse.Data.Latin, asmaulHusnaResponse.Data.Arab, asmaulHusnaResponse.Data.Arti), logger)
		return
	}

	asmaulHusnaResponse := getAsmaulHusnaByLatin(command, ASMAUL_HUSNA_API_URL, s, m, logger)

	if asmaulHusnaResponse.Data.Urutan == 0 {
		utils.MessageWithReply(s, m, "Asmaul Husna tidak ditemukan", logger)
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", asmaulHusnaResponse.Data.Urutan, asmaulHusnaResponse.Data.Latin, asmaulHusnaResponse.Data.Arab, asmaulHusnaResponse.Data.Arti), logger)
}

func loopAsmaulHusnaMessage(start int, end int, asmaulHusnaResponse []entities.AsmaulHusna, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	content := ``

	for _, v := range asmaulHusnaResponse[start:end] {
		content += fmt.Sprintf("%d - %s - %s - %s\n", v.Urutan, v.Latin, v.Arab, v.Arti)
	}

	utils.MessageWithReply(s, m, content, logger)
}

func getAllAsmaulHusna(ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) []entities.AsmaulHusna {
	response, err := http.Get(ASMAUL_HUSNA_API_URL + "/api/all")
	if err != nil {
		utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading Asma'ul Husna", logger)
		return nil
	}

	var asmaulHusnaResponse entities.AsmaulHusnaResponse
	err = json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asmaul Husna", logger)
		return nil
	}

	return asmaulHusnaResponse.Data
}

func getAsmaulHusnaByUrutan(urutan int, ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) entities.AsmaulHusnaByLatinOrUrutanResponse {
	response, err := http.Get(ASMAUL_HUSNA_API_URL + "/api/" + strconv.Itoa(urutan))
	if err != nil {
		utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
		return entities.AsmaulHusnaByLatinOrUrutanResponse{}
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading Asmaul Husna", logger)
		return entities.AsmaulHusnaByLatinOrUrutanResponse{}
	}

	var asmaulHusnaResponse entities.AsmaulHusnaByLatinOrUrutanResponse

	err = json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asma'ul Husna", logger)
		return entities.AsmaulHusnaByLatinOrUrutanResponse{}
	}

	return asmaulHusnaResponse
}

func getAsmaulHusnaByLatin(latin string, ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) entities.AsmaulHusnaByLatinOrUrutanResponse {
	response, err := http.Get(ASMAUL_HUSNA_API_URL + "/api/latin/" + latin)
	if err != nil {
		utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
		return entities.AsmaulHusnaByLatinOrUrutanResponse{}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading Asmaul Husna", logger)
		return entities.AsmaulHusnaByLatinOrUrutanResponse{}
	}

	var asmaulHusnaResponse entities.AsmaulHusnaByLatinOrUrutanResponse

	err = json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asmaul Husna", logger)
		return entities.AsmaulHusnaByLatinOrUrutanResponse{}
	}

	return asmaulHusnaResponse
}
