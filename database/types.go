package database

import (
	"database/sql"
	"time"
)

type Dialect string

const (
    PostgresDialect Dialect = "postgres"
    OracleDialect  Dialect = "oracle"
)

type Config struct {
    Dialect         Dialect
    Host            string
    Port            int
    Username        string
    Password        string
    Database        string
    Schema          string        // For PostgreSQL
    MaxOpenConns    int
    MaxIdleConns    int
    ConnMaxLifetime time.Duration
    SSLMode         string        // For PostgreSQL
    ConnectTimeout  int          // Connection timeout in seconds
    ServiceName     string
}

type DB struct {
    *sql.DB
    config *Config
}