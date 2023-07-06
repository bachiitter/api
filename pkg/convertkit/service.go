package convertkit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func AddSubscriber(email string) (AddSubscriberResponse, error) {
	accessToken := os.Getenv("CONVERTKIT_API_KEY")
	formId := os.Getenv("CONVERTKIT_FORM_ID")

	url := fmt.Sprintf("https://api.convertkit.com/v3/forms/%s/subscribe?api_key=%s", formId, accessToken)

	req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf(`{"email": "%s"}`, email)))
	if err != nil {
		return AddSubscriberResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return AddSubscriberResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AddSubscriberResponse{}, err
	}

	var data AddSubscriberResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return AddSubscriberResponse{}, err
	}

	return data, err
}

func GetTotalSubs() (TotalsSubsResponse, error) {
	accessToken := os.Getenv("CONVERTKIT_API_SECRET")
	formId := os.Getenv("CONVERTKIT_FORM_ID")

	url := fmt.Sprintf("https://api.convertkit.com/v3/forms/%s/subscriptions?api_secret=%s", formId, accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TotalsSubsResponse{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return TotalsSubsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TotalsSubsResponse{}, err
	}

	var data TotalsSubsResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return TotalsSubsResponse{}, err
	}

	return data, err
}
