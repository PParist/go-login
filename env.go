package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getENV(c *fiber.Ctx) error {
	_SECRECT := os.Getenv("SECRECT")
	if _SECRECT != "" {
		return c.JSON(fiber.Map{
			"SECRECT": os.Getenv("SECRECT"),
		})
	}
	return c.JSON(fiber.Map{
		"SECRECT": "default secrect",
	})
}
