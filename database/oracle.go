package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func buildOracleConnString(config *Config) string {
    /* Creates a Connection String for oracle connection */
    connStr := fmt.Sprintf(`user="%s" password="%s" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%d))(CONNECT_DATA=(SERVICE_NAME=%s)(SERVER=DEDICATED)(PROGRAM=%s)))"`,
        config.Username,
        config.Password,
        config.Host,
        config.Port,
        config.Database,
        config.ServiceName,
    )

    return connStr
}

func connectOracle(config *Config) (*sql.DB, error) {
	/* 
		This function is used to Create and check the oracle database connection.
	*/

    connStr := buildOracleConnString(config)
    connection, err := sql.Open("godror", connStr)
    if err != nil {
        return nil, fmt.Errorf("error opening oracle connection: %w", err)
    }

    connection.SetMaxOpenConns(config.MaxOpenConns)
    connection.SetMaxIdleConns(config.MaxIdleConns)
    connection.SetConnMaxLifetime(config.ConnMaxLifetime)

    if err := connection.Ping(); err != nil {
        connection.Close()
        return nil, fmt.Errorf("error connecting to oracle: %w", err)
    }

    return connection, nil
}