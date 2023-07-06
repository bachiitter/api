package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetAuthHeader() string {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	return "Basic " + base64.RawStdEncoding.Strict().EncodeToString([]byte(clientId+":"+clientSecret)) + "="
}

func GetAccessToken() (TokenResponse, error) {
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token?grant_type=refresh_token&refresh_token="+refreshToken, nil)
	if err != nil {
		return TokenResponse{}, err
	}

	req.Header.Set("Authorization", GetAuthHeader())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TokenResponse{}, err
	}

	var data TokenResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return TokenResponse{}, err
	}

	return data, nil
}

func Client(query string) ([]byte, error) {
	tokens, err := GetAccessToken()
	if err != nil {
		return []byte{}, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1%s", query)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Authorization", "Bearer "+tokens.AccessToken)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return body, err
}
