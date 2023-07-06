package unsplash

import "time"

type PhotosResponse struct {
	ID             string    `json:"id"`
	Slug           string    `json:"slug"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	Color          string    `json:"color"`
	BlurHash       string    `json:"blur_hash"`
	Description    *string   `json:"description"`
	AltDescription *string   `json:"alt_description"`
	Urls           Urls      `json:"urls"`
	Links          Links     `json:"links"`
	User           User      `json:"user"`
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
	SmallS3 string `json:"small_s3"`
}

type Links struct {
	Self     string `json:"self"`
	Html     string `json:"html"`
	Download string `json:"download"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
}

type Downloads struct {
	Total int `json:"total"`
}

type Views struct {
	Total int `json:"total"`
}

type StatsResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Downloads Downloads `json:"downloads"`
	Views     Views     `json:"views"`
}
