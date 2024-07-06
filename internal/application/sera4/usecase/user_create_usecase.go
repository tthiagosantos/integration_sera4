package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"integrations_apis/internal/application/sera4/dto"
	"net/http"
	"os"
)

type UserCreateUsecase struct {
	httpClient *http.Client
}

func NewCreateUserUseCase() *UserCreateUsecase {
	return &UserCreateUsecase{
		httpClient: &http.Client{},
	}
}

func (uc *UserCreateUsecase) Execute(url, session string, user dto.UserRequest) (string, error) {
	request := map[string]interface{}{
		"user": map[string]interface{}{
			"first_name":        user.FirstName,
			"last_name":         user.LastName,
			"user_group":        "user",
			"email":             user.Email,
			"language":          "en",
			"device_restricted": true,
		},
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tws-membership-id", os.Getenv("SERA4_MEMBERSHIP"))
	req.Header.Set("tws-organization-token", os.Getenv("SERA4_ORGANIZATION_TOKEN"))
	req.Header.Set("Authorization", "Bearer "+session)

	res, err := uc.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("status: %s", res.Status)
	}

	var membership map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&membership)
	if err != nil {
		return "", err
	}

	data, ok := membership["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to decode data response: %w", err)
	}

	id, ok := data["membership_id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to decode membership_id response: %w", err)
	}
	return id, nil
}
