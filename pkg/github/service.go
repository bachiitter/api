package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Client(query string) ([]byte, error) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	url := fmt.Sprintf("https://api.github.com%s", query)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

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

func GetRepoCommits(owner string, repo string, page string) ([]CommitData, error) {

	url := fmt.Sprintf("/repos/%s/%s/commits?page=%s", owner, repo, page)

	body, err := Client(url)
	if err != nil {
		return []CommitData{}, err
	}

	var data []CommitData

	if err := json.Unmarshal(body, &data); err != nil {
		return []CommitData{}, err
	}

	return data, err
}

func GetRepo(owner string, repo string) (Repository, error) {

	url := fmt.Sprintf("/repos/%s/%s", owner, repo)

	body, err := Client(url)
	if err != nil {
		return Repository{}, err
	}

	var data Repository

	if err := json.Unmarshal(body, &data); err != nil {
		return Repository{}, err
	}

	return data, err
}

func GetUser(user string) (User, error) {

	url := fmt.Sprintf("/users/%s", user)

	body, err := Client(url)
	if err != nil {
		return User{}, err
	}

	var data User

	if err := json.Unmarshal(body, &data); err != nil {
		return User{}, err
	}

	return data, err
}
