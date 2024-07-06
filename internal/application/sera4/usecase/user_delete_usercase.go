package usecase

import (
	"io"
	"net/http"
	"os"
)

type UserDeleteUseCase struct {
	httpClient *http.Client
}

func NewDeleteUserCase() *UserDeleteUseCase {
	return &UserDeleteUseCase{
		httpClient: &http.Client{},
	}
}

func (uc *UserDeleteUseCase) Execute(session, url string) error {
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

	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
