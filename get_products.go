package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getProducts(c *fiber.Ctx) error {
	defer timeTrack(time.Now(), "getProducts")
	products, err := getProductsDetails()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan productgegevens niet ophalen",
		})
	}

	return c.Render("products", fiber.Map{
		"Products": products,
	})
}

func getProductsDetails() ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, title, images, descr, specs, price, weight, length, width, height FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Images, &p.Descr, &p.Specs, &p.Price, &p.Weight, &p.Length, &p.Width, &p.Height); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
