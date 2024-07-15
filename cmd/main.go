package main

import (
	"errors"
	Bot "jokes_bot/internal/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const DELAY float32 = 24.00

func getEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New(".env файл не найден")
	}

	if api, exists := os.LookupEnv("API_KEY"); !exists {
		return "", errors.New("записи об api ключе не существует")

	} else {
		return api, nil
	}

}

func main() {

	log.Println("Начало работы")
	log.Println("получение данных из .env файла")

	api_key, err := getEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("api ключ (" + api_key + ") получен.")

	bot := Bot.NewBot(api_key)
	log.Println("бот инициализирован")

	if resp, err := bot.UploadJoke(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("новая шутейка:", resp)
	}

	// TODO: бот должен работать регулярно по DELAY
	// for {
	// }

}
