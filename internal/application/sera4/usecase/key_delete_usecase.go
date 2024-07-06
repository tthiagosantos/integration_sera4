package usecase

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type KeyDeleteUseCase struct {
	httpClient *http.Client
}

func NewDeleteKeyUseCase() *KeyDeleteUseCase {
	return &KeyDeleteUseCase{
		httpClient: &http.Client{},
	}
}

func (uc *KeyDeleteUseCase) Execute(session, url string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tws-membership-id", os.Getenv("SERA4_MEMBERSHIP"))
	req.Header.Set("tws-organization-token", os.Getenv("SERA4_ORGANIZATION_TOKEN"))
	req.Header.Set("Authorization", "Bearer "+session)

	resp, err := uc.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	log.Println("STATUS: ", resp.Status)
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	return errors.New(resp.Status)
}
