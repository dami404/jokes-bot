package parser

import (
	json "encoding/json"
	"errors"
	"io"
	"jokes_bot/internal/utils"
	"net/http"
	"strings"
)

// type JokesSources struct {
// 	vk *VKPublic
// 	tg *TGChannel
// }

type VKPublic struct {
	Owner_id     string
	Domain       string
	Access_token string
	Date         int
}

type TGChannel struct {
	// TODO: more fields
	Date int
}

type Parser interface {
	SendRequest() (joke string, err error)
}

func NewParser(owner_id string, domain string, access_token string) *VKPublic {
	return &VKPublic{
		Owner_id:     owner_id,
		Domain:       domain,
		Access_token: access_token,
		Date:         0,
	}
}

// TODO: TG PARSER
func (tg TGChannel) NewParser() *TGChannel {
	return &TGChannel{
		Date: 0,
	}
}

func (vk *VKPublic) SendRequest() (string, error) {
	req := "https://api.vk.ru/method/wall.get?v=5.199&owner_id=" + vk.Owner_id + "&domain=" + vk.Domain
	res, err := http.NewRequest(http.MethodGet, req, nil)
	utils.CheckErrors(err)
	res.Header.Add("Authorization", "Bearer "+vk.Access_token)
	resp, err := http.DefaultClient.Do(res)
	utils.CheckErrors(err)
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	utils.CheckErrors(err)
	return string(body), nil
}

// TODO: refactor
func (tg *TGChannel) SendRequest() (string, error) {
	return "", nil
}

func validateJoke(joke string) (string, error) {
	isValidated := strings.Contains(joke, "http") || strings.Contains(joke, "t.me")
	utils.ThrowErrorsIfFalse(isValidated, errors.New("Это реклама"))
	return joke, nil
}

// TODO: убрать ненужные переносы, кроме '/n -' (прямая речь)
// TODO: \" - обработать экранирование
// func formatJoke(joke string) string {
// }

func convertToJSON(jsonString *string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(*jsonString), &data)
	return data, err
}

func GetJoke(p Parser, date *int) (string, error) {
	var text string
	resp, err := p.SendRequest()
	utils.CheckErrors(err)

	data, err := convertToJSON(&resp)
	utils.CheckErrors(err)
	response, ok := data["response"].(map[string]interface{})
	utils.ThrowErrorsIfFalse(ok, errors.New("Тег response отсутствует"))
	items, ok := response["items"].([]interface{})
	utils.ThrowErrorsIfFalse(ok, errors.New("Тег items отсутствует"))
	jokePost := items[1].(map[string]interface{})
	*date, _ = jokePost["date"].(int)
	text = jokePost["text"].(string)
	// joke, err := validateJoke(text)
	// return joke, err
	return text, err
}
