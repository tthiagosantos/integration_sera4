package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SessionCreateUsecase struct {
	httpClient *http.Client
}

func NewSessionUseCase() *SessionCreateUsecase {
	return &SessionCreateUsecase{
		httpClient: http.DefaultClient,
	}
}

func (uc *SessionCreateUsecase) Execute(url string) (string, error) {
	request := map[string]string{
		"username": os.Getenv("SERA4_USERNAME"),
		"password": os.Getenv("SERA4_PASSWORD"),
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to execute request, status code: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	twsToken, ok := response["tws_token"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to decode tws_token response")
	}
	twsTokenData, ok := twsToken["tws_token_data"].(string)
	if !ok {
		return "", fmt.Errorf("failed to decode tws_token_data response")
	}
	return twsTokenData, nil
}
