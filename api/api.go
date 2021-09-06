package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var Precision int = 1
var App *fiber.App
var Api fiber.Router

func GetPrecision(c *fiber.Ctx) error {
	s := fmt.Sprintf("{'precision': %d}", Precision)
	return c.Send([]byte(s))

}

func AssignPrecision(c *fiber.Ctx) error {
	newPrecision := c.Query("newPrecision")
	fmt.Println(newPrecision)
	Precision, _ = strconv.Atoi(newPrecision)
	return c.Send([]byte("{'message': 'success'}"))

}
