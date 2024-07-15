package bot

import (
	"io"
	"net/http"
)

type Bot struct {
	API_KEY string
}

type BotAbilities interface {
	NewBot(api_key string) *Bot
	SendRequest() (string, error)
	UploadJoke() (string, error)
	getJoke() string
}

func NewBot(api_key string) *Bot {
	return &Bot{api_key}
}

// TODO: refactor
func (bot *Bot) SendRequest(method string) (string, error) {
	if resp, err := http.Get("https://api.telegram.org/bot" + bot.API_KEY + "/" + method); err != nil {
		return "", err
	} else {
		if body, err := io.ReadAll(resp.Body); err != nil {
			resp.Body.Close()
			return "", err
		} else {
			resp.Body.Close()
			return string(body), nil
		}

	}

}

// TODO: шутку брать из парсера
func (bot *Bot) UploadJoke() (string, error) {
	joke := "Штирлиц играл в карты и проигрался. Но Штирлиц умел делать хорошую мину при плохой игре. Когда Штирлиц покинул компанию, мина сработала."
	method := "sendMessage?chat_id=@white_rock_off&text=" + joke
	body, err := bot.SendRequest(method)
	return body, err

}
