package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func deleteProduct(c *fiber.Ctx) error {

	// ID uit URL parameter halen
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Als ID geen integer is, geef een foutmelding terug
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ongeldig product ID",
		})
	}

	// if err := c.BodyParser(&product); err != nil {
	// 	fmt.Println(err)
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Kan productgegevens niet verwerken",
	// 	})
	// }

	_, err = db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan product niet verwijderen uit database",
		})
	}

	fmt.Println("id", id, "verwjderd")
	return c.JSON(fiber.Map{
		"message": "Product succesvol verwijderd",
	})
}
