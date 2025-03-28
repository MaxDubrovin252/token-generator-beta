package service

import "token-generator/pkg/reposiroty"

type Token interface {
}

type Service struct {
	Token
}

func NewService(repos *reposiroty.Repository) *Service {
	return &Service{}
}
