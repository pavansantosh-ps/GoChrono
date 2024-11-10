package database

import (
	"database/sql"
	"fmt"
)

func New(config *Config) (*DB, error) {
    if config == nil {
        return nil, ErrNilConfig
    }

    var (
        connection  *sql.DB
        err error
    )

    switch config.Dialect {
    case PostgresDialect:
        connection, err = connectPostgres(config)
    case OracleDialect:
        connection, err = connectOracle(config)
    default:
        return nil, fmt.Errorf("%w: %s", ErrInvalidDialect, config.Dialect)
    }

    if err != nil {
        return nil, err
    }

    return &DB{
        DB:     connection,
        config: config,
    }, nil
}

func (connection *DB) Close() error {
    if connection.DB == nil {
        return ErrNoConnection
    }
    return connection.DB.Close()
}

func (connection *DB) Ping() error {
    if connection.DB == nil {
        return ErrNoConnection
    }
    return connection.DB.Ping()
}

func (connection *DB) GetDialect() Dialect {
    return connection.config.Dialect
}