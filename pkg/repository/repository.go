package repository

import (
	"github.com/jmoiron/sqlx"
)

type TokenOP interface {
	CreateToken(message, token string) (int, error)
	GetMessage(token string) (string, error)
}

type Repository struct {
	TokenOP
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TokenOP: NewPostgresCreate(db),
	}
}
