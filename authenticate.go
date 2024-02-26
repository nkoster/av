package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func authenticate(c *fiber.Ctx) error {
    sess, err := store.Get(c)

    if err != nil {
        fmt.Println(err)
    }
    // Controleer of de sessie een gebruiker bevat
    if sess.Get("email") == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Niet geautoriseerd"})
    }

    return c.Next()
}
