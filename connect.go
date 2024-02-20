package main

import (
	"database/sql"
	"fmt"
	"os"
)

func connect() {
	dbName := os.Getenv("PG_DB")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASS")
	socketDir := "/var/run/postgresql"
	connStr := fmt.Sprintf("postgres://%s:%s@/%s?host=%s", user, password, dbName, socketDir)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Fout bij het verbinden met de database:", err)
	}
}
