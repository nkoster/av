package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	UrlTitle string  `json:"url_title"`
	Image    string  `json:"image"`
	Descr    string  `json:"descr"`
	Specs    string  `json:"specs"`
	Price    int     `json:"price"`
	Weight   int     `json:"weight"`
}

func getProducts(c *fiber.Ctx) error {
	defer timeTrack(time.Now(), "getProducts")
	products, err := getProductsDetails()
	if err != nil {
		fmt.Println(err)
		// Terugsturen van een serverfout als het ophalen van de productgegevens mislukt
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan productgegevens niet ophalen",
		})
	}

	return c.JSON(fiber.Map{
		"products": products,
	})
}

func getProductsDetails() ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, title, url_title, image, descr, specs, price, weight FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.UrlTitle, &p.Image, &p.Descr, &p.Specs, &p.Price, &p.Weight); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
