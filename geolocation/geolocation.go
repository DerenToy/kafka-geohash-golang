package main

import (
	"encoding/json"
	"fmt"
	"kafka-golang/api"
	"kafka-golang/producer"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	api.App = fiber.New()
	api.Api = api.App.Group("/api/v1")

	api.Api.Post("/geolocation", createGeolocation)
	api.Api.Get("/precision", GetPrecision)
	api.Api.Post("/precision", AssignPrecision)
	api.App.Listen(":3000")

}
func GetPrecision(c *fiber.Ctx) error {
	s := fmt.Sprintf("{'precision': %d}", api.Precision)
	return c.Send([]byte(s))

}

func AssignPrecision(c *fiber.Ctx) error {
	newPrecision := c.Query("newPrecision")
	fmt.Println(newPrecision)
	api.Precision, _ = strconv.Atoi(newPrecision)
	return c.Send([]byte("{'message': 'success'}"))

}

func createGeolocation(c *fiber.Ctx) error {

	// Instantiate new Message struct
	cmt := new(producer.Geolocation)

	//  Parse body into comment struct
	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	// convert body into bytes and send it to kafka
	cmtInBytes, err := json.Marshal(cmt)
	producer.PushToQueue("geolocations", cmtInBytes)

	// Return Comment in JSON format
	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}

	return err
}
