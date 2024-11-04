package bot

import (
	"io"
	"jokes_bot/internal/utils"
	"log"
	"net/http"
)

type Bot struct {
	API_KEY string
}

type BotAbilities interface {
	NewBot(api_key string) *Bot
	UploadJoke() (string, error)
	getJoke() string
}

func NewBot(api_key string) *Bot {
	return &Bot{api_key}
}

func (bot *Bot) sendRequest(method string) (string, error) {
	resp, err := http.Get("https://api.telegram.org/bot" + bot.API_KEY + "/" + method)
	utils.CheckErrors(err)

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	utils.CheckErrors(err)
	return string(body), nil
}

// TODO: шутку брать из парсера
func (bot *Bot) UploadJoke(joke string) (string, error) {
	method := "sendMessage?chat_id=@white_rock_off&text=" + joke
	body, err := bot.sendRequest(method)
	log.Println("новая шутейка:", joke)
	return body, err
}
