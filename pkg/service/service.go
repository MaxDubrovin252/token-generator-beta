package service

import (
	"token-generator/pkg/repository"
)

type TokenOP interface {
	CreateToken(message, token string) (int, error)
	GetMessage(token string) (string, error)
}

type Service struct {
	TokenOP
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TokenOP: NewCreateService(repos.TokenOP),
	}
}
