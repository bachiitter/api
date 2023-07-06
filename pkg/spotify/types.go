package spotify

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type SpotifyImage struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type SpotifyArtist struct {
	Name   string         `json:"name"`
	Images []SpotifyImage `json:"images"`
}

type SpotifyAlbum struct {
	Images []SpotifyImage `json:"images"`
}

type SpotifyItem struct {
	Name         string `json:"name"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Album      SpotifyAlbum    `json:"album"`
	Artists    []SpotifyArtist `json:"artists"`
	Duration   int             `json:"duration_ms"`
	PreviewUrl string          `json:"preview_url"`
}

type PlayerResponse struct {
	IsPlaying bool `json:"is_playing"`
	Device    struct {
		Id               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
	} `json:"device"`
	Item       SpotifyItem `json:"item"`
	ProgressMS int         `json:"progress_ms"`
}

type NowPlayingResponse struct {
	IsPlaying bool        `json:"is_playing"`
	Item      SpotifyItem `json:"item"`
}

type LastPlayedResponse struct {
	Items []struct {
		Track SpotifyItem `json:"track"`
	} `json:"items"`
}

type TopTracksReponse struct {
	Items []SpotifyItem `json:"items"`
}
