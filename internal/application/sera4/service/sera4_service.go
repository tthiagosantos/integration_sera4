package service

import (
	"integrations_apis/internal/application/sera4/dto"
	"integrations_apis/internal/application/sera4/usecase"
	"log"
	"net/http"
	"os"
)

type Sera4Service struct {
	baseURL    string
	httpClient *http.Client
}

func NewService() *Sera4Service {
	return &Sera4Service{
		baseURL:    os.Getenv("SERA4_BASE_URL"),
		httpClient: &http.Client{},
	}
}

func (s *Sera4Service) CreateUser(user dto.UserRequest) (string, error) {
	session, err := s.CreatedSession()
	if err != nil {
		return "", err
	}
	url := s.baseURL + "users"
	result, err := usecase.NewCreateUserUseCase().Execute(url, session, user)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *Sera4Service) DeleteUser(id string) error {
	url := s.baseURL + "users/" + id
	session, err := s.CreatedSession()
	if err != nil {
		return err
	}
	err = usecase.NewDeleteUserCase().Execute(session, url)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sera4Service) GetUser(id string) (map[string]interface{}, error) {
	url := s.baseURL + "users/" + id
	session, err := s.CreatedSession()
	if err != nil {
		return nil, err
	}
	user, err := usecase.NewUserUseCase().Execute(url, session)
	if err != nil {
		return nil, err
	}
	log.Println(id)
	log.Println(user)
	return user, nil
}

func (s *Sera4Service) CreatedSession() (string, error) {
	session, err := usecase.NewSessionUseCase().Execute(s.baseURL + "sessions")
	if err != nil {
		return "", err
	}
	return session, nil
}

func (s *Sera4Service) CreateKey(key dto.KeyDTO) (int, error) {
	url := s.baseURL + "keys"
	session, err := s.CreatedSession()
	if err != nil {
		return 0, err
	}
	ok, err := usecase.NewKeyCreate().Execute(url, session, key)
	if err != nil {
		return 0, err
	}
	return ok, nil
}

func (s *Sera4Service) DeleteKey(id string) error {
	url := s.baseURL + "keys/" + id
	session, err := s.CreatedSession()
	if err != nil {
		return err
	}
	err = usecase.NewDeleteKeyUseCase().Execute(session, url)
	if err != nil {
		return err
	}
	return nil
}
