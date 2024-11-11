package database

import (
	"errors"
)

var (
	ErrInvalidDialect = errors.New("invalid database dialect")
	ErrNoConnection   = errors.New("no database connection established")
	ErrNilConfig      = errors.New("database configuration is nil")
)
