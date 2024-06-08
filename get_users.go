package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getUsers(c *fiber.Ctx) error {
	defer timeTrack(time.Now(), "getUsers")
	users, err := getUsersDetails()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get users.",
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func getUsersDetails() ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var p User
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
