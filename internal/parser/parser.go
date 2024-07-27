package parser

import (
	json "encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type JokesSources struct {
	ws *WebSite
	vk *VKPublic
	tg *TGChannel
}

type WebSite struct {
	Url     string
	From_me bool
}

type VKPublic struct {
	Owner_id     string
	Domain       string
	Access_token string
	From_me      bool
	Date         int
}

type TGChannel struct {
	// TODO: more fields
	From_me bool
	Date    int
}

type Parser interface {
	SendRequest() (joke string)
}

func NewParser(owner_id, domain, access_token string) *JokesSources {
	return &JokesSources{
		vk: &VKPublic{
			Owner_id:     owner_id,
			Domain:       domain,
			Access_token: access_token,
			From_me:      true,
			Date:         0,
		},
		tg: &TGChannel{
			From_me: false,
			Date:    0,
		},
		ws: &WebSite{
			Url:     "https://www.trees-and-lambdas.info/matushansky/stirlitz.html",
			From_me: false,
		},
	}
}

// TODO: TG PARSER

// TODO: WEBSITE PARSER

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

// TODO: refactor 123-127
func (js *JokesSources) GetJoke() (string, error) {
	var resp string
	var err error
	var date *int
	var text string
	switch true {
	case js.vk.From_me:
		resp, err = js.vk.SendRequest()
		if err != nil {
			return "", err
		}
		date = &js.vk.Date
	case js.tg.From_me:
		return "", nil
	case js.ws.From_me:
		return "", nil
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
