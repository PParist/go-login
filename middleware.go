package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// UserData represents the user data extracted from the JWT token
type UserData struct {
	Email string
	Role  string
}

// Middleware to extract user data from JWT
func validateToken(c *fiber.Ctx) error {
	fmt.Println("restricted")

	token := c.Locals("user").(*jwt.Token)
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	user := new(UserData)
	claims := token.Claims.(jwt.MapClaims)
	user.Email = claims["email"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, user)

	return c.Next()
}

func isAdmin(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)

	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized role is not admin"})
		//return fiber.ErrUnauthorized.Message()
	}

	return c.Next()
}
