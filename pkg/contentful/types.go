package contentful

type PostsSlugsResponse struct {
	Data struct {
		ProjectCollection struct {
			Items []struct {
				Slug string `json:"slug"`
			} `json:"items"`
		} `json:"postCollection"`
	} `json:"data"`
}
type PostsResponse struct {
	Data struct {
		ProjectCollection struct {
			Items []struct {
				Sys struct {
					ID             string `json:"id"`
					PublishedAt    string `json:"publishedAt"`
					FirstPublished string `json:"firstPublishedAt"`
				} `json:"sys"`
				Title       string    `json:"title"`
				Description string    `json:"description"`
				Thumbnail   Thumbnail `json:"thumbnail"`
				Slug        string    `json:"slug"`
				Featured    bool      `json:"featured"`
			} `json:"items"`
		} `json:"postCollection"`
	} `json:"data"`
}

type PostResponse struct {
	Data struct {
		ProjectCollection struct {
			Items []struct {
				Sys struct {
					ID             string `json:"id"`
					PublishedAt    string `json:"publishedAt"`
					FirstPublished string `json:"firstPublishedAt"`
				} `json:"sys"`
				Title       string    `json:"title"`
				Description string    `json:"description"`
				Thumbnail   Thumbnail `json:"thumbnail"`
				Slug        string    `json:"slug"`
				Featured    bool      `json:"featured"`
				Content     string    `json:"content"`
			} `json:"items"`
		} `json:"postCollection"`
	} `json:"data"`
}

type Thumbnail struct {
	URL         string `json:"url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Description string `json:"description"`
}

type ProjectsResponse struct {
	Data struct {
		ProjectCollection struct {
			Items []struct {
				Sys struct {
					ID             string `json:"id"`
					PublishedAt    string `json:"publishedAt"`
					FirstPublished string `json:"firstPublishedAt"`
				} `json:"sys"`
				Title       string    `json:"title"`
				Description string    `json:"description"`
				Link        string    `json:"link"`
				Stack       []string  `json:"stack"`
				Thumbnail   Thumbnail `json:"thumbnail"`
			} `json:"items"`
		} `json:"projectCollection"`
	} `json:"data"`
}
