package reposiroty

import "github.com/jmoiron/sqlx"

type Token interface {
}

type Repository struct {
	Token
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
