package storage

import "errors"

var (
	ErrTokenNotFound = errors.New("cannot find your token")
	ErrTokenExists   = errors.New("your token is already exists")
)
