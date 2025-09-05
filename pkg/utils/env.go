package utils

import (
	"katou-megumi/pkg/configs"
	"log"

	"github.com/spf13/viper"
)

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
	v := configs.NewViper()

	_ = v.BindEnv("DATABASE_URL")
	_ = v.BindEnv("DISCORD_TOKEN")
	_ = v.BindEnv("REMOVE_BG_API_KEY")
	_ = v.BindEnv("REMOVE_BG_API_URL")
	_ = v.BindEnv("JOKES_API_URL")
	_ = v.BindEnv("ANIME_QUOTE_API_URL")
	_ = v.BindEnv("DISTRO_INFO_API_URL")
	_ = v.BindEnv("DOA_API_URL")
	_ = v.BindEnv("QURAN_API_URL")
	_ = v.BindEnv("IMAGE_API_URL")
	_ = v.BindEnv("GEMINI_API_KEY")
	_ = v.BindEnv("ASMAUL_HUSNA_API_URL")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config file: %v", err)
		}
	}

	return EnvVariables{
		DISCORD_TOKEN:        v.GetString("DISCORD_TOKEN"),
		REMOVE_BG_API_KEY:    v.GetString("REMOVE_BG_API_KEY"),
		REMOVE_BG_API_URL:    v.GetString("REMOVE_BG_API_URL"),
		JOKES_API_URL:        v.GetString("JOKES_API_URL"),
		ANIME_QUOTE_API_URL:  v.GetString("ANIME_QUOTE_API_URL"),
		DISTRO_INFO_API_URL:  v.GetString("DISTRO_INFO_API_URL"),
		DOA_API_URL:          v.GetString("DOA_API_URL"),
		QURAN_API_URL:        v.GetString("QURAN_API_URL"),
		IMAGE_API_URL:        v.GetString("IMAGE_API_URL"),
		GEMINI_API_KEY:       v.GetString("GEMINI_API_KEY"),
		ASMAUL_HUSNA_API_URL: v.GetString("ASMAUL_HUSNA_API_URL"),
	}
}
