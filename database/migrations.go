package database

import (
	"context"
	"database/sql"
	"fmt"
)

func (connection *DB) CreateScheduleTable() (sql.Result, error) {
    statement := "CREATE TABLE IF NOT EXISTS  %s (%s, %s)"
    v := []interface{}{"Tasks", "Task_ID varchar(30) UNIQUE", "Schedule varchar(100)"}
    return connection.Exec(fmt.Sprintf(statement, v...))
}

func (connection *DB) Setup() (error){
	ctx := context.Background()
    tx, err := connection.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

    _, err = connection.CreateScheduleTable()

	if err != nil {
		return err
	}

    tx.Commit()

	return nil
}