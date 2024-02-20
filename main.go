package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB
var err error

func main() {

	// var titles []string

	app := fiber.New()

	if err = godotenv.Load(); err != nil {
		fmt.Println("No .env file in current directory.")
	}
	
	connect()

	app.Get("/products/titles", getProducts)
	// app.Get("/products/titles", func(c *fiber.Ctx) error {
	// 	titles, err = getProductTitles()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		// Terugsturen van een serverfout als het ophalen van de producttitels mislukt
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": "Kan producttitels niet ophalen",
	// 		})
	// 	}

	// 	return c.JSON(fiber.Map{
	// 		"titles": titles,
	// 	})
	// })

	defer db.Close()
	log.Fatal(app.Listen(":3000"))
}
