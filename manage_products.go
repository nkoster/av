package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func manageProducts(c *fiber.Ctx) error {
	defer timeTrack(time.Now(), "getProducts")
	products, err := getProductsDetails()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan productgegevens niet ophalen",
		})
	}

	return c.Render("manageproducts", fiber.Map{
		"Products": products,
	})
}
