package main

import "github.com/gofiber/fiber/v2"

func authenticate(c *fiber.Ctx) error {
    sess := store.Get(c)

    // Controleer of de sessie een gebruiker bevat
    if sess.Get("email") == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Niet geautoriseerd"})
    }

    return c.Next()
}
