package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

type EnvVariables struct {
	DISCORD_TOKEN        string
	REMOVE_BG_API_KEY    string
	REMOVE_BG_API_URL    string
	JOKES_API_URL        string
	ANIME_QUOTE_API_URL  string
	DISTRO_INFO_API_URL  string
	DOA_API_URL          string
	QURAN_API_URL        string
	IMAGE_API_URL        string
	GEMINI_API_KEY       string
	ASMAUL_HUSNA_API_URL string
}

func Env() EnvVariables {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return EnvVariables{
		DISCORD_TOKEN:        os.Getenv("DISCORD_TOKEN"),
		REMOVE_BG_API_KEY:    os.Getenv("REMOVE_BG_API_KEY"),
		REMOVE_BG_API_URL:    os.Getenv("REMOVE_BG_API_URL"),
		JOKES_API_URL:        os.Getenv("JOKES_API_URL"),
		ANIME_QUOTE_API_URL:  os.Getenv("ANIME_QUOTE_API_URL"),
		DISTRO_INFO_API_URL:  os.Getenv("DISTRO_INFO_API_URL"),
		DOA_API_URL:          os.Getenv("DOA_API_URL"),
		QURAN_API_URL:        os.Getenv("QURAN_API_URL"),
		IMAGE_API_URL:        os.Getenv("IMAGE_API_URL"),
		GEMINI_API_KEY:       os.Getenv("GEMINI_API_KEY"),
		ASMAUL_HUSNA_API_URL: os.Getenv("ASMAUL_HUSNA_API_URL"),
	}
}
