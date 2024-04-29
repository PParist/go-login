package main

import (
	"github.com/gofiber/fiber/v2"
)

func uploadFile(c *fiber.Ctx) error {
	_file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = c.SaveFile(_file, "./upload/"+_file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("upload complete !!")
}
