package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.new"

	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	stmt, err := db.Prepare(`
	
	CREATE TABLE IF NOT EXISTS tokens(
	id INTEGER PRIMARY KEY,
	token TEXT NOT NULL UNIQUE,
	message TEXT NIT NULL);
	CREATE INDEX IF NOT EXISTS idx_token ON tokens(token)
	
	
	
	);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	_, err = stmt.Exec()

	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return &Storage{db: db}, nil
}
func (s *Storage) NewToken(message string, token string) (int64, error) {
	const op = "sqlite.NewToken"

	stmt, err := s.db.Prepare(`INSERT INTO tokens(message, token) VALUES(?,?)`)

	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	res, err := stmt.Exec(message, token)

	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}
	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	return id, nil
}
