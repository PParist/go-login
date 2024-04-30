package main

import (
	"log"
	"os"

	_ "github.com/PParist/go-login/docs"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	swagger "github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var users []User // Slice to store users
const userContextKey = "user"

// @title Login API
// @description This is a sample server for a Login API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	app.Get("/swagger/*", swagger.HandlerDefault)
	users = append(users, User{ID: 1, Title: "1984", Author: "George Orwell"})
	users = append(users, User{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	return c.JSON(users)
	// })

	app.Post("/login", userLogin)

	// JWT Middleware & Middleware check role(restricted)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRECT"))},
	}), validateToken)

	adminGroup := app.Group("/admin")
	userGroup := app.Group("/user")
	adminGroup.Use(isAdmin)

	//Set Route
	adminGroup.Get("/users", getUsers)
	adminGroup.Get("/users/:id", getUserById)
	//render template
	userGroup.Get("/template", renderTemplate)
	//get env
	adminGroup.Get("/env", getENV)

	adminGroup.Post("/createusers", createUser)
	adminGroup.Post("/upload", uploadFile)

	adminGroup.Put("/updateusers/:id", updateUser)
	adminGroup.Delete("/removeuser/:id", removeUserById)

	app.Listen(":8080")
}
