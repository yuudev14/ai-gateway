package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Settings = &SettingsData{}
)

type SettingsData struct {
	LLM_BASE_URL string
}

func Setup() {
	godotenv.Load("./.env")
	SetEnv()
}

func SetEnv() {
	Settings.LLM_BASE_URL = os.Getenv("LLM_BASE_URL")
}
