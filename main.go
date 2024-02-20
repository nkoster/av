package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	app := fiber.New()

	if err = godotenv.Load(); err != nil {
		log.Fatal("No .env file in current directory.")
	}
	
	connect()

	app.Get("/products/titles", getProducts)

	defer db.Close()
	log.Fatal(app.Listen(":3000"))
}
