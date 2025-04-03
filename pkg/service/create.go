package service

import (
	"token-generator/pkg/repository"
)

type CreateService struct {
	repo repository.TokenOP
}

func NewCreateService(repo repository.TokenOP) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateToken(message, token string) (int, error) {
	return s.repo.CreateToken(message, token)
}
func (s *CreateService) GetMessage(token string) (string, error) {
	return s.repo.GetMessage(token)
}
