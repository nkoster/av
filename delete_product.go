package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func deleteProduct(c *fiber.Ctx) error {

	// Get ID from URL parameter
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Error when product ID != integer
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	_, err = db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Removing product failed!",
		})
	}

	fmt.Println("id", id, "removed")
	return c.SendString("<script>window.location.reload();</script>")

}
