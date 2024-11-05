package bot

import (
	"io"
	"jokes_bot/internal/utils"
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

func (bot *Bot) sendRequest(joke string) (string, error) {
	url := "https://api.telegram.org/bot" + bot.API_KEY + "/sendMessage?chat_id=@white_rock_off&text=" + joke
	resp, err := http.Get(url)
	utils.CheckErrors("bot-sendRequest-1-", err)

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	utils.CheckErrors("bot-sendRequest-2-", err)
	return string(body), nil
}

// TODO: шутку брать из парсера
func (bot *Bot) UploadJoke(joke string) (string, error) {
	body, err := bot.sendRequest(joke)
	return body, err
}
