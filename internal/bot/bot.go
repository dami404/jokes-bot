package bot

import (
	"io"
	"jokes_bot/internal/parser"
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

// TODO: refactor
func (bot *Bot) sendRequest(method string) (string, error) {
	if resp, err := http.Get("https://api.telegram.org/bot" + bot.API_KEY + "/" + method); err != nil {
		return "", err
	} else {
		if body, err := io.ReadAll(resp.Body); err != nil {
			defer resp.Body.Close()
			return "", err
		} else {
			resp.Body.Close()
			return string(body), nil
		}

	}

}

// TODO: шутку брать из парсера
func (bot *Bot) UploadJoke(js *parser.JokesSources) (string, error) {
	if joke, err := js.GetJoke(); err != nil {
		return "", err
	} else {
		method := "sendMessage?chat_id=@white_rock_off&text=" + joke
		body, err := bot.sendRequest(method)
		log.Println("новая шутейка:", joke)

		return body, err
	}
}
