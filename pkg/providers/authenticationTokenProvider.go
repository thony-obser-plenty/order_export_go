package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

func FetchToken(url, apiKey string) (*Token, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ApiKey", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tkn Token
	err = json.Unmarshal(body, &tkn)
	if err != nil {
		return nil, err
	}

	return &tkn, nil
}
