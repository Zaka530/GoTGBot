package system

import (
	"github.com/joho/godotenv"
	"os"
)

func BotToken() string {
	// Use the absolute path to your .env file
	err := godotenv.Load("/Users/kamron/GolandProjects/TgBot/.env")

	if err != nil {
		panic(err)
	}
	return os.Getenv("BOT_TOKEN")
}
