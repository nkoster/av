package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getUsers(c *fiber.Ctx) error {
	defer timeTrack(time.Now(), "getUsers")
	users, err := getUsersDetails()
	if err != nil {
		fmt.Println(err)
		// Terugsturen van een serverfout als het ophalen van de productgegevens mislukt
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kan productgegevens niet ophalen",
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func getUsersDetails() ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p User
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
