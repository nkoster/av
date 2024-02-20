package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func newProduct(c *fiber.Ctx) error {

	defer timeTrack(time.Now(), "newProduct")
	var p Product

	// Parse de JSON body naar de Product struct
	if err := c.BodyParser(&p); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Kan productgegevens niet verwerken",
		})
	}

	// Voeg het nieuwe product toe aan de database
	_, err := db.Exec("INSERT INTO products (title, url_title, image, descr, specs, price, weight) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		p.Title, p.UrlTitle, p.Image, p.Descr, p.Specs, p.Price, p.Weight)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan product niet toevoegen aan database",
		})
	}

	// Stuur een succesbericht terug
	return c.JSON(fiber.Map{
		"message": "Product succesvol toegevoegd",
	})
}
