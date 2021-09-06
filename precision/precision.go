package precision

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/api/v1/precision", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.Send(c.Body())

	})

}
