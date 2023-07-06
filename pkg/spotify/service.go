package spotify

import (
	"encoding/json"
	"fmt"
)

func GetPlayer() (PlayerResponse, error) {

	url := "/me/player"

	body, err := Client(url)
	if err != nil {
		return PlayerResponse{}, err
	}

	var data PlayerResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return PlayerResponse{}, err
	}

	return data, err
}

func GetNowPlaying() (NowPlayingResponse, error) {

	url := "/me/player/currently-playing"

	body, err := Client(url)
	if err != nil {
		return NowPlayingResponse{}, err
	}

	var data NowPlayingResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return NowPlayingResponse{}, err
	}

	return data, err
}

func GetLastPlayed(limit string) (LastPlayedResponse, error) {
	if limit == "" || limit == "'" {
		limit = "1"
	}

	url := fmt.Sprintf("/me/player/recently-played?limit=%s", limit)

	body, err := Client(url)
	if err != nil {
		return LastPlayedResponse{}, err
	}

	var data LastPlayedResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return LastPlayedResponse{}, err
	}

	return data, err
}

func GetTopTracks(timeRange string) (TopTracksReponse, error) {
	if timeRange == "" || timeRange == "'" {
		timeRange = "medium_term"
	}

	url := fmt.Sprintf("/me/top/tracks?limit=10&time_range=%s", timeRange)

	body, err := Client(url)
	if err != nil {
		return TopTracksReponse{}, err
	}

	var data TopTracksReponse

	if err := json.Unmarshal(body, &data); err != nil {
		return TopTracksReponse{}, err
	}

	return data, err
}
