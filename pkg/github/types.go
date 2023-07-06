package github

import "time"

type Committer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type Commit struct {
	Author    Author    `json:"author"`
	Committer Committer `json:"committer"`
	Message   string    `json:"message"`
	URL       string    `json:"url"`
}

type CommitData struct {
	Sha       string `json:"sha"`
	Commit    Commit `json:"commit"`
	HTMLURL   string `json:"html_url"`
	Author    Owner  `json:"author"`
	Committer Owner  `json:"committer"`
}

type Owner struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}

type Repository struct {
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Private          bool        `json:"private"`
	Owner            Owner       `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      string      `json:"description"`
	StargazersCount  int         `json:"stargazers_count"`
	Language         string      `json:"language"`
	License          interface{} `json:"license"`
	Watchers         int         `json:"watchers"`
	NetworkCount     int         `json:"network_count"`
	SubscribersCount int         `json:"subscribers_count"`
}

type User struct {
	Login                   string    `json:"login"`
	ID                      int       `json:"id"`
	AvatarURL               string    `json:"avatar_url"`
	HTMLURL                 string    `json:"html_url"`
	Name                    string    `json:"name"`
	Company                 string    `json:"company"`
	Blog                    string    `json:"blog"`
	Location                string    `json:"location"`
	Email                   string    `json:"email"`
	Hireable                bool      `json:"hireable"`
	Bio                     string    `json:"bio"`
	TwitterUsername         string    `json:"twitter_username"`
	PublicRepos             int       `json:"public_repos"`
	Followers               int       `json:"followers"`
	Following               int       `json:"following"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	TwoFactorAuthentication bool      `json:"two_factor_authentication"`
}
