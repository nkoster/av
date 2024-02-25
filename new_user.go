package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func newUser(c *fiber.Ctx) error {

	var user User

	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kan verzoek niet verwerken"})
	}

	// Het wachtwoord hashen
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kan wachtwoord niet hashen"})
	}

	// Nieuwe gebruiker toevoegen aan de database
	_, err = db.Exec("INSERT INTO users (first_name, last_name, username, password, email, phone, street_name, house_number, city, postal_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", user.FirstName, user.LastName, user.Username, string(hashedPassword), user.Email, user.Phone, user.StreetName, user.HouseNumber, user.City, user.PostalCode)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kan gebruiker niet aanmaken"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Gebruiker succesvol aangemaakt"})
}
