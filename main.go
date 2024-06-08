package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var DEBUG int
var store = session.New()

func main() {

	if err = godotenv.Load(); err != nil {
		log.Fatal("No .env file in current directory.")
	}

	connect()

	if DEBUG, err = strconv.Atoi(os.Getenv("DEBUG")); err != nil {
		DEBUG = 0
	}

	engine := html.New("./templates", ".html")
	engine.Reload(true)
	if DEBUG == 1 {
		engine.Debug(true)
	}

	engine.AddFunc("mod", func(a, b int) int {
		return a % b
	})
	engine.AddFunc("next", func(i int) int {
		return i + 1
	})
	engine.AddFunc("images", func(input string) []string {
		return strings.Split(input, ",")
	})
	engine.AddFunc("splitLines", func(s string) []string {
		return strings.Split(s, "\n")
	})
	engine.AddFunc("split", strings.Split)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	UI := os.Getenv("UI")
	fmt.Printf("Serving static files: \t%s\n", UI)

	app.Static("/", UI)
	if DEBUG == 1 {
		app.Use("/", func(c *fiber.Ctx) error {
			startTime := time.Now()
			err := c.Next()
			fmt.Printf("Request to %s, Method: %s, Status: %d, Time: %v\n", c.OriginalURL(), c.Method(), c.Response().StatusCode(), time.Since(startTime))
			return err
		})
	}

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
	app.Get("/productjson/:id", getProductJSON)
	app.Get("/users", authenticate, getUsers)
	app.Post("/api/newproduct", newProduct)
	app.Post("/api/uploadimages", uploadImages)
	app.Get("/api/deleteproduct/:id", deleteProduct)
	app.Get("/api/manageproducts", manageProducts)
	app.Post("/api/updateproduct", updateProduct)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))

}
