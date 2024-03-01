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

	fmt.Printf("update_product.go: %+v\n\n", p)

	// Werk het product bij in de database
	_, err := db.Exec("UPDATE products SET title = $1, images = $2, descr = $3, specs = $4, price = $5, weight = $6, length = $7, width = $8, height = $9 WHERE id = $10",
		p.Title, p.Images, p.Descr, p.Specs, p.Price, p.Weight, p.Length, p.Width, p.Height, p.ID)
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
