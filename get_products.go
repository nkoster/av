package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getProducts(c *fiber.Ctx) error {
	titles, err := getProductTitles()
	if err != nil {
		fmt.Println(err)
		// Terugsturen van een serverfout als het ophalen van de producttitels mislukt
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan producttitels niet ophalen",
		})
	}

	return c.JSON(fiber.Map{
		"titles": titles,
	})
}

func getProductTitles() ([]string, error) {

	var titles []string
	rows, err := db.Query("SELECT title FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return titles, nil
}
