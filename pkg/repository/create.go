package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CreatePostgres struct {
	db *sqlx.DB
}

func NewPostgresCreate(db *sqlx.DB) *CreatePostgres {
	return &CreatePostgres{db: db}
}

const tokenTable = "tokens"

func (r *CreatePostgres) CreateToken(message, token string) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(message,token) VALUES($1,$2)", tokenTable)

	row := r.db.QueryRow(query, message, token)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *CreatePostgres) GetMessage(token string) (string, error) {

	var message string

	query := fmt.Sprintf("SELECT message FROM %s WHERE token=$1", tokenTable)

	err := r.db.QueryRow(query, token).Scan(&message)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("cannot get a token%s", err.Error())
		}

		return "", err
	}

	return message, nil

}
