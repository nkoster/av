package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func login(c *fiber.Ctx) error {
    type LoginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    var request LoginRequest
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kan verzoek niet verwerken"})
    }

    // Query de database voor de gebruiker
    var hashedPassword string
    err := db.QueryRow("SELECT password FROM users WHERE email = $1", request.Email).Scan(&hashedPassword)
    if err != nil {
        // Gebruiker niet gevonden of andere fout
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Onjuiste email of wachtwoord"})
    }

    // Vergelijk het opgegeven wachtwoord met het gehashte wachtwoord in de database
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password))
    if err != nil {
        // Fout betekent dat het wachtwoord niet overeenkomt
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Onjuiste gebruikersnaam of wachtwoord"})
    }

    // Wachtwoord is correct, maak een sessie
    sess := store.Get(c)

    // Sla de gebruikersinformatie op in de sessie
    sess.Set("email", request.Email)
    if err := sess.Save(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kan sessie niet opslaan"})
    }

    return c.SendStatus(fiber.StatusOK)
}
