package parser

import (
	json "encoding/json"
	"errors"
	"io"
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

// TODO: refactor
func (vk *VKPublic) SendRequest() (string, error) {
	req := "https://api.vk.ru/method/wall.get?v=5.199&owner_id=" + vk.Owner_id + "&domain=" + vk.Domain
	res, err := http.NewRequest(http.MethodGet, req, nil)
	if err != nil {
		return "", err
	} else {
		res.Header.Add("Authorization", "Bearer "+vk.Access_token)
		resp, err := http.DefaultClient.Do(res)
		if err != nil {
			return "", err
		}
		if body, err := io.ReadAll(resp.Body); err != nil {
			defer resp.Body.Close()
			return "", err
		} else {
			resp.Body.Close()
			return string(body), nil
		}
	}
}

// TODO: refactor
func (tg *TGChannel) SendRequest() (string, error) {
	return "", nil
}

// TODO: убрать ненужные переносы, кроме '/n -' (прямая речь)
func validateJoke(joke string) (string, error) {
	if strings.Contains(joke, "http") || strings.Contains(joke, "t.me") {
		return "", errors.New("это не шутка, а реклама")
	}
	return joke, nil
}

func convertToJSON(jsonString *string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(*jsonString), &data)
	return data, err
}

// TODO: refactor 108-112
func GetJoke(p Parser, date *int) (string, error) {
	// var date *int
	var text string
	resp, err := p.SendRequest()
	if err != nil {
		return "", err
	}

	// println(resp)
	data, err := convertToJSON(&resp)
	if err != nil {
		return "", err
	}
	if response, ok := data["response"].(map[string]interface{}); ok {
		if items, ok := response["items"].([]interface{}); ok {
			jokePost := items[1].(map[string]interface{})
			*date, _ = jokePost["date"].(int)
			text = jokePost["text"].(string)
		}
	}
	joke, err := validateJoke(text)
	return joke, err
}
