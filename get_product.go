package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getProduct(c *fiber.Ctx) error {
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

	return c.Render("product", fiber.Map{
		"Product": product,
	})
}

func getProductDetails(id int) (*Product, error) {
	var product Product

	// Pas de query aan om PostgreSQL's $1 placeholder te gebruiken voor de parameter
	row := db.QueryRow("SELECT id, title, url_title, images, descr, specs, price, weight, length, width, height FROM products WHERE id = $1", id)
	if err := row.Scan(&product.ID, &product.Title, &product.UrlTitle, &product.Images, &product.Descr, &product.Specs, &product.Price, &product.Weight, &product.Length, &product.Width, &product.Height); err != nil {
		return nil, err
	}

	return &product, nil
}

