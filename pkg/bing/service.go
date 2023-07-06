package bing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetTime(city string) (TimeResponse, error) {
	accessToken := os.Getenv("BING_MAPS_API_TOKEN")

	url := fmt.Sprintf("https://dev.virtualearth.net/REST/v1/TimeZone/?query=%s&key=%s", city, accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TimeResponse{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return TimeResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TimeResponse{}, err
	}

	var data TimeResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return TimeResponse{}, err
	}

	return data, err
}

func GetLocation(point string) (LocationResponse, error) {
	accessToken := os.Getenv("BING_MAPS_API_TOKEN")

	url := fmt.Sprintf("https://dev.virtualearth.net/REST/v1/Locations/%s?key=%s", point, accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	var data LocationResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return LocationResponse{}, err
	}

	return data, err
}
