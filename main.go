package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	UrlTitle string  `json:"url_title"`
	Image    string  `json:"image"`
	Descr    string  `json:"descr"`
	Specs    string  `json:"specs"`
	Price    int     `json:"price"`
	Weight   int     `json:"weight"`
}

var db *sql.DB
var err error

func main() {

	app := fiber.New()

	if err = godotenv.Load(); err != nil {
		log.Fatal("No .env file in current directory.")
	}
	
	connect()

	UI := os.Getenv("UI")
	fmt.Printf("Serving static files: \t%s\n", UI)
	app.Static("/", UI)

	app.Get("/products", getProducts)

	app.Post("/newproduct", newProduct)

	defer db.Close()
	log.Fatal(app.Listen(":3000"))
}
