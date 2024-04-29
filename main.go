package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var users []User // Slice to store users

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize standard Go html template engine
	engine := html.New("./view", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	users = append(users, User{ID: 1, Title: "1984", Author: "George Orwell"})
	users = append(users, User{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	return c.JSON(users)
	// })
	app.Post("/login", userLogin)

	//MiddleWare check role
	app.Use(restricted)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRECT")),
	}))
	//Set Route
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUserById)
	//render template
	app.Get("/template", renderTemplate)
	//get env
	app.Get("/env", getENV)

	app.Post("/createusers", createUser)
	app.Post("/upload", uploadFile)

	app.Put("/updateusers/:id", updateUser)
	app.Delete("/removeuser/:id", removeUserById)

	app.Listen(":8080")
}
