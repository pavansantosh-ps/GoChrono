package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pavansantosh-ps/gochrono/database"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: .env file not found: %v", err)
    }

    dbConfig, err := database.NewConfig()
    if err != nil {
        log.Fatalf("Failed to create database config: %v", err)
    }

    connection, err := database.New(dbConfig)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer connection.Close()

    if err := connection.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Printf("Successfully connected to %s database", connection.GetDialect())

}

