package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/session/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

var store = session.New()

func main() {

	app := fiber.New()

	if err = godotenv.Load(); err != nil {
		log.Fatal("No .env file in current directory.")
	}

	connect()

	UI := os.Getenv("UI")
	fmt.Printf("Serving static files: \t%s\n", UI)

	app.Static("/", UI)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use("/api", authenticate)

	app.Post("/login", login)

	app.Get("/admin", authenticate, func(c *fiber.Ctx) error {
		return c.SendString("Je bent ingelogd en hebt toegang tot deze beveiligde route!")
	})

	app.Post("/newuser", newUser)
	app.Get("/products", getProducts)
	app.Get("/users", authenticate, getUsers)
	app.Post("/api/newproduct", newProduct)
	app.Post("/api/updateproduct", updateProduct)

	defer db.Close()

	log.Fatal(app.Listen(":3000"))

}
