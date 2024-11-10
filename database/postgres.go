package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func buildPostgresConnString(config *Config) string {
	/* Creates a Connection String for Postgres connection */
    sslMode := config.SSLMode
    if sslMode == "" {
        sslMode = "disable"
    }

    connStr := fmt.Sprintf("application_name=%s host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.ServiceName,
        config.Host,
        config.Port,
        config.Username,
        config.Password,
        config.Database,
        sslMode,
    )

    if config.Schema != "" {
        connStr += fmt.Sprintf(" search_path=%s", config.Schema)
    }

    if config.ConnectTimeout > 0 {
        connStr += fmt.Sprintf(" connect_timeout=%d", config.ConnectTimeout)
    }

    return connStr
}

func connectPostgres(config *Config) (*sql.DB, error) {
	/* 
		This function is used to Create and check the postgres database connection.
	*/

    connStr := buildPostgresConnString(config)
    connection, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("error opening postgres connection: %w", err)
    }

    connection.SetMaxOpenConns(config.MaxOpenConns)
    connection.SetMaxIdleConns(config.MaxIdleConns)
    connection.SetConnMaxLifetime(config.ConnMaxLifetime)

    if err := connection.Ping(); err != nil {
        connection.Close()
        return nil, fmt.Errorf("error connecting to postgres: %w", err)
    }

    return connection, nil
}