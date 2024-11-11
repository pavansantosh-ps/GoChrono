package database

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func NewConfig() (*Config, error) {
	dialect := Dialect(os.Getenv("DB_DIALECT"))
	if dialect != PostgresDialect && dialect != OracleDialect {
		return nil, fmt.Errorf("%w: %s", ErrInvalidDialect, dialect)
	}

	connName := os.Getenv("SERVICE_NAME")
	if connName == "" {
		connName = "Unknown"
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = defaultPort(dialect)
	}

	maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if maxOpenConns == 0 {
		maxOpenConns = 20
	}

	maxIdleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if maxIdleConns == 0 {
		maxIdleConns = 5
	}

	connMaxLifetime, _ := time.ParseDuration(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if connMaxLifetime == 0 {
		connMaxLifetime = 5 * time.Minute
	}

	connectTimeout, _ := strconv.Atoi(os.Getenv("DB_CONNECT_TIMEOUT"))
	if connectTimeout == 0 {
		connectTimeout = 10
	}

	return &Config{
		Dialect:         dialect,
		Host:            os.Getenv("DB_HOST"),
		Port:            port,
		Username:        os.Getenv("DB_USERNAME"),
		Password:        os.Getenv("DB_PASSWORD"),
		Database:        os.Getenv("DB_DATABASE"),
		Schema:          os.Getenv("DB_SCHEMA"),
		MaxOpenConns:    maxOpenConns,
		MaxIdleConns:    maxIdleConns,
		ConnMaxLifetime: connMaxLifetime,
		SSLMode:         os.Getenv("DB_SSL_MODE"),
		ConnectTimeout:  connectTimeout,
		ServiceName:     connName,
	}, nil
}

func defaultPort(dialect Dialect) int {
	switch dialect {
	case PostgresDialect:
		return 5432
	case OracleDialect:
		return 1521
	default:
		return 0
	}
}
