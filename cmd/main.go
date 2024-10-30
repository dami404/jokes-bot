package main

import (
	"errors"
	Bot "jokes_bot/internal/bot"
	"jokes_bot/internal/parser"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const DELAY float64 = 3.00

// TODO: refactor
func getEnv() (api_key, access_token, user_id, domain string, err error) {
	if err := godotenv.Load(); err != nil {
		return "", "", "", "", errors.New(".env файл не найден")
	}

	if api_key, exists := os.LookupEnv("API_KEY"); !exists {
		return "", "", "", "", errors.New("записи об api ключе не существует")
	} else if access_token, exists := os.LookupEnv("ACCESS_TOKEN"); !exists {
		return "", "", "", "", errors.New("записи об access token не существует")
	} else if user_id, exists := os.LookupEnv("USER_ID"); !exists {
		return "", "", "", "", errors.New("записи об user id не существует")
	} else if domain, exists := os.LookupEnv("VK_DOMAIN"); !exists {
		return "", "", "", "", errors.New("записи об domain не существует")
	} else {
		return api_key, access_token, user_id, domain, nil

	}

}

func main() {

	log.Println("Начало работы")
	log.Println("получение данных из .env файла")

	api_key, access_token, user_id, domain, err := getEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("api ключ (" + api_key + ") получен.")
	log.Println("access token (" + access_token + ") получен.")
	log.Println("user id  (" + user_id + ") получен.")
	log.Println("domain (" + domain + ") получен.")

	// tg := parser.TGChannel{
	// 	Date: 0,
	// }

	vkParser := parser.NewParser(user_id, domain, access_token)
	bot := Bot.NewBot(api_key)
	log.Println("бот инициализирован")

	for {
		joke, err := parser.GetJoke(vkParser, &vkParser.Date)
		if err != nil {
			log.Println(err)
		}

		resp, err := bot.UploadJoke(joke)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("новая шутейка:", resp)
		}

		time.Sleep(time.Hour * time.Duration(DELAY))
	}
}
