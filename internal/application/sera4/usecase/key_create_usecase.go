package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"integrations_apis/internal/application/sera4/dto"
	"log"
	"net/http"
	"os"
)

type KeyCreateUsecase struct {
	httpClient *http.Client
}

func NewKeyCreate() *KeyCreateUsecase {
	return &KeyCreateUsecase{
		httpClient: &http.Client{},
	}
}

func (uc *KeyCreateUsecase) Execute(url, session string, key dto.KeyDTO) (int, error) {
	payload, err := json.Marshal(map[string]interface{}{"key": key})
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tws-membership-id", os.Getenv("SERA4_MEMBERSHIP"))
	req.Header.Set("tws-organization-token", os.Getenv("SERA4_ORGANIZATION_TOKEN"))
	req.Header.Set("Authorization", "Bearer "+session)

	resp, err := uc.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to execute request, status code: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, fmt.Errorf("failed to execute request: %w", err)
	}
	data, ok := response["data"].([]interface{})
	if !ok || len(data) == 0 {
		return 0, errors.New("invalid data response")
	}
	log.Println(response["data"])
	log.Println("======")
	log.Println(data)
	firstItem, ok := data[0].(map[string]interface{})
	if !ok {
		return 0, errors.New("invalid data item")
	}

	id, ok := firstItem["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("failed to decode id response: %v", firstItem)
	}

	return int(id), nil
}
