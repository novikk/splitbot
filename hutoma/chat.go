package hutoma

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const HUTOMA_BASE_URL = "https://api.hutoma.ai"

type HutomaClient struct {
	BotID     string
	DevKey    string
	ClientKey string
}

func (c *HutomaClient) Chat(query string) (hutomaChatResponse, error) {
	query = url.QueryEscape(query) // prepare query

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/ai/%s/chat?q=%s", HUTOMA_BASE_URL, c.BotID, query), nil)
	if err != nil {
		return hutomaChatResponse{}, errors.New("failed creating get chat request: " + err.Error())
	}

	client := http.Client{}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.DevKey))

	res, err := client.Do(req)
	if err != nil {
		return hutomaChatResponse{}, errors.New("failed performing get chat request: " + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return hutomaChatResponse{}, errors.New("failed reading body from response: " + err.Error())
	}

	var chatRes hutomaChatResponse
	err = json.Unmarshal(body, &chatRes)
	if err != nil {
		fmt.Println(string(body))
		return hutomaChatResponse{}, errors.New("failed unmarshaling json response: " + err.Error())
	}

	return chatRes, nil
}
