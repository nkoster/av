package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func connect() {
	dbName := os.Getenv("PG_DB")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASS")
	socketDir := "/var/run/postgresql"
	connStr := fmt.Sprintf("postgres://%s:%s@/%s?host=%s", user, password, dbName, socketDir)
	if db, err = sql.Open("postgres", connStr); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}
