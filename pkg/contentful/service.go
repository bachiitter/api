package contentful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Client(query string) ([]byte, error) {
	accessToken := os.Getenv("CONTENTFUL_ACCESS_TOKEN")
	spaceId := os.Getenv("CONTENTFUL_SPACE_ID")

	jsonMapInstance := map[string]string{
		"query": query,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)
	url := fmt.Sprintf("https://graphql.contentful.com/content/v1/spaces/%s", spaceId)

	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonResult))

	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

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

func GetAllProjects() (ProjectsResponse, error) {

	body, err := Client(ProjectsQuery)
	if err != nil {
		return ProjectsResponse{}, err
	}

	var data ProjectsResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return ProjectsResponse{}, err
	}

	return data, err
}

func GetAllPosts() (PostsResponse, error) {

	body, err := Client(PostsQuery)
	if err != nil {
		return PostsResponse{}, err
	}

	var data PostsResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return PostsResponse{}, err
	}

	return data, err
}

func GetPost(slug string) (PostResponse, error) {
	query := fmt.Sprintf(`{
		postCollection(where: {slug: "%s"}) {
			items {
			 sys {
					id
					publishedAt
					firstPublishedAt
				}
				title
				description
				thumbnail {
					url
					width
					height
					description
				}
				slug
				featured
				content
			}
		}
	}
	`, slug)

	body, err := Client(query)
	if err != nil {
		return PostResponse{}, err
	}

	var data PostResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return PostResponse{}, err
	}

	return data, err
}
