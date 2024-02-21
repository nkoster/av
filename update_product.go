package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func updateProduct(c *fiber.Ctx) error {

	defer timeTrack(time.Now(), "updateProduct")

	var p Product

	// Parse de JSON body naar de Product struct
	if err := c.BodyParser(&p); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Kan productgegevens niet verwerken",
		})
	}

	// Controleer of het product ID is meegegeven
	if p.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is vereist",
		})
	}

	// Werk het product bij in de database
	_, err := db.Exec("UPDATE products SET title = $1, url_title = $2, images = $3, descr = $4, specs = $5, price = $6, weight = $7, length = $8, width = $9, height = $10 WHERE id = $11",
		p.Title, p.UrlTitle, p.Images, p.Descr, p.Specs, p.Price, p.Weight, p.Length, p.Width, p.Height, p.ID)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan product niet bijwerken in database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product succesvol bijgewerkt",
	})
}
