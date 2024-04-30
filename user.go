package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Book struct to hold book data
type User struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type LoginUser struct {
	//ID       int    `json:"id"`
	//Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var loginUser = LoginUser{
	Email:    "user@example.com",
	Password: "password123",
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUserById(c *fiber.Ctx) error {
	_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if _id != 0 { // ตรวจสอบว่า ID ไม่เป็น 0 ก่อนที่จะค้นหา
		newusers := []User{}
		for _, v := range users {
			if v.ID == _id {
				newusers = append(newusers, v)
			}
		}
		if len(newusers) > 0 {
			return c.Status(fiber.StatusOK).JSON(newusers)
		} else {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}
	} else {
		return c.Status(fiber.StatusNotFound).SendString("Not Found id 0")
	}
}

func createUser(c *fiber.Ctx) error {
	newBook := new(User)
	//TODO: pars body to struct
	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	users = append(users, *newBook)
	return c.Status(fiber.StatusCreated).JSON(users)
}

func updateUser(c *fiber.Ctx) error {
	_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	newUser := new(User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		for i, user := range users {
			if user.ID == _id {
				users[i].Title = newUser.Title
				users[i].Author = newUser.Author
				return c.Status(fiber.StatusOK).JSON(users[i])
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}
}

func removeUserById(c *fiber.Ctx) error {

	_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	for i, user := range users {
		if user.ID == _id {
			//TODO: remove
			//[1,2,3,4,5]
			//[1,2] + [4,5] = [1,2,4,5]
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(fiber.StatusOK)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func userLogin(c *fiber.Ctx) error {
	_newUser := new(LoginUser)
	if err := c.BodyParser(_newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//check authorize
	if _newUser.Email != loginUser.Email || _newUser.Password != loginUser.Password {
		return fiber.ErrUnauthorized
	}
	/*//วิธีแรก
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = _newUser.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	_token, err := token.SignedString([]byte(os.Getenv("JWT_SECRECT")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}*/

	// Create the Claims
	claims := jwt.MapClaims{
		"email": _newUser.Email,
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRECT")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"data": _newUser, "token": t, "resault": "login success"})
}
