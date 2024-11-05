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

const DELAY float64 = 0.00

var ENVs = map[string]string{
	"API_KEY":      "",
	"ACCESS_TOKEN": "",
	"USER_ID":      "",
	"VK_DOMAIN":    "",
}

var err error

func ENVInit() {
	err := godotenv.Load()
	utils.CheckErrors("ENVInit-", err)
}

// TODO: refactor
func getENV(parName string) (value string, err error) {
	err = godotenv.Load()
	utils.CheckErrors("getENV-", err)

	value, exists := os.LookupEnv(parName)
	utils.ThrowErrorsIfFalse("getENV", exists, errors.New("Такой переменной ENV не существует"))
	return value, nil
}

func main() {

	log.Println("Начало работы")
	log.Println("Получение данных из .env файла")

	ENVInit()
	for k := range ENVs {
		ENVs[k], err = getENV(k)
		utils.CheckErrors("main-1-", err)
	}

	vkParser := parser.NewParser(ENVs["USER_ID"], ENVs["VK_DOMAIN"], ENVs["ACCESS_TOKEN"])
	bot := Bot.NewBot(ENVs["API_KEY"])
	log.Println("Бот инициализирован")

	oldJoke := ""

	for {
		time.Sleep(0)

		joke, err := parser.GetJoke(vkParser)
		utils.CheckErrors("main-2-", err)

		if oldJoke == joke {
			log.Println("Анекдот уже был")
			continue
		}
		oldJoke = joke

		_, err = bot.UploadJoke(joke)
		utils.CheckErrors("main-3-", err)

		log.Println("Новый анекдот!")

	}
}
