package main

import (
	"errors"
	Bot "jokes_bot/internal/bot"
	"jokes_bot/internal/parser"
	"jokes_bot/internal/utils"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const DELAY float64 = 3.00

var ENVs = map[string]string{
	"API_KEY":      "",
	"ACCESS_TOKEN": "",
	"USER_ID":      "",
	"VK_DOMAIN":    "",
}

var err error

func ENVInit() {
	err := godotenv.Load()
	utils.CheckErrors(err)
}

// TODO: refactor
func getENV(parName string) (value string, err error) {
	err = godotenv.Load()
	value, exists := os.LookupEnv(parName)
	utils.ThrowErrorsIfFalse(exists, errors.New("Такой переменной ENV не существует"))
	return value, nil
}

func main() {

	log.Println("Начало работы")
	log.Println("Получение данных из .env файла")

	ENVInit()
	for k, _ := range ENVs {
		ENVs[k], err = getENV(k)
		utils.CheckErrors(err)
	}
	log.Println("api ключ получен.")
	log.Println("access token получен.")
	log.Println("user id  получен.")
	log.Println("domain получен.")

	// tg := parser.TGChannel{
	// 	Date: 0,
	// }

	vkParser := parser.NewParser(ENVs["USER_ID"], ENVs["VK_DOMAIN"], ENVs["ACCESS_TOKEN"])
	bot := Bot.NewBot(ENVs["API_KEY"])
	log.Println("Бот инициализирован")

	for {
		joke, err := parser.GetJoke(vkParser, &vkParser.Date)
		resp, err := bot.UploadJoke(joke)
		if err != nil {
			log.Println(err)
		}

		log.Println("новая шутейка:", resp)
		time.Sleep(time.Hour * time.Duration(DELAY))
	}
}
