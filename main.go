package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

var store = session.New()

func main() {

	engine := html.New("./templates", ".html")
	engine.Reload(true)
	engine.Debug(true)

	engine.AddFunc("mod", func(a, b int) int {
        return a % b
    })
	engine.AddFunc("next", func(i int) int {
		return i + 1
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

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

	app.Get("/validatesession", validateSession)
	app.Post("/newuser", newUser)
	app.Get("/products", getProducts)
	app.Get("/product/:id", getProduct)
	app.Get("/users", authenticate, getUsers)
	app.Post("/api/newproduct", newProduct)
	app.Get("/api/deleteproduct/:id", deleteProduct)
	app.Get("/api/manageproducts", manageProducts)
	app.Post("/api/updateproduct", updateProduct)

	defer db.Close()

	log.Fatal(app.Listen(":3000"))

}
