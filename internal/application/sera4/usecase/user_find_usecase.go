package usecase

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type UserQueryUsecase struct {
	httpClient *http.Client
}

func NewUserUseCase() *UserQueryUsecase {
	return &UserQueryUsecase{
		httpClient: &http.Client{},
	}
}

func (uc *UserQueryUsecase) Execute(url, session string) (map[string]interface{}, error) {
	log.Println("URL: " + url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tws-membership-id", os.Getenv("SERA4_MEMBERSHIP"))
	req.Header.Set("tws-organization-token", os.Getenv("SERA4_ORGANIZATION_TOKEN"))
	req.Header.Set("Authorization", "Bearer "+session)

	res, err := uc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	var userData map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&userData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
