package main

import (
	"github.com/gofiber/fiber/v2"
)

func renderTemplate(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"Title": "Hello,World"})
}
