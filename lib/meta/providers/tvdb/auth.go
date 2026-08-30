package tvdb

import (
	"animeta/lib/core/env"
	"animeta/lib/core/fetch"
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
)

func login() (string, error) {
	apiKey, err := env.GetVar("TVDB_APIKEY")
	if err != nil {
		return "", err
	}

	body, _ := json.Marshal(map[string]string{
		"apikey": apiKey,
	})

	resp, err := http.Post(BaseURL+"/login", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("tvdb login request error: %w", err)
	}

	var loginResp LoginResponse

	fetch.ExtractResponseJsonBody(resp, &loginResp)
	if err != nil {
		return "", fmt.Errorf("tvdb login body parsing error: %w", err)
	}

	return loginResp.Data.Token, nil
}

func Authorize() (Client, error) {
	token, err := login()

	if err != nil {
		return Client{}, err
	}

	return Client{
		Token: token,
	}, nil
}
