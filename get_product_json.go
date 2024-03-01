package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getProductJSON(c *fiber.Ctx) error {
	defer timeTrack(time.Now(), "getProduct")
	// ID uit URL parameter halen
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Als ID geen integer is, geef een foutmelding terug
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ongeldig product ID",
		})
	}

	product, err := getProductDetails(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan productgegevens niet ophalen",
		})
	}

	fmt.Printf("getProductJSON.go: %+v\n\n", product)
	// Geef de productgegevens terug als JSON
	return c.JSON(product)
}
