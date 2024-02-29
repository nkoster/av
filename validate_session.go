package main

import (
	"github.com/gofiber/fiber/v2"
)

func validateSession(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil || sess.Fresh() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Ongeldige sessie",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
