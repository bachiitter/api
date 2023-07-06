package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetUserPhotos(username string) ([]PhotosResponse, error) {
	accessToken := os.Getenv("UNSPLASH_API_KEY")

	url := fmt.Sprintf("https://api.unsplash.com/users/%s/photos", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []PhotosResponse{}, err
	}

	req.Header.Set("Authorization", "Client-ID "+accessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []PhotosResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []PhotosResponse{}, err
	}

	var data []PhotosResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return []PhotosResponse{}, err
	}

	return data, err
}

func GetUserStats(username string) (StatsResponse, error) {
	accessToken := os.Getenv("UNSPLASH_API_KEY")

	url := fmt.Sprintf("https://api.unsplash.com/users/%s/statistics", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return StatsResponse{}, err
	}

	req.Header.Set("Authorization", "Client-ID "+accessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return StatsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return StatsResponse{}, err
	}

	var data StatsResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return StatsResponse{}, err
	}

	return data, err
}
