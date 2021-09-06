package router

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App, channel chan int) {
	app.Post("/api/v1?precision=", func(c *fiber.Ctx) {
		var precision int
		fmt.Println("çaloştı mı")
		precision, _ = strconv.Atoi(c.Body())
		channel <- precision
		fmt.Println("çalıştı")

	})

}
