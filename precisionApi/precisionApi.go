package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/geohash", handler)
	//http.Handle()
	http.ListenAndServe(":8080", nil)

	// app := fiber.New()
	// //precision.SetupRoutes(app)
	// go router.SetupRoutes(app, cc.Channel)
	// app.Listen(":3001")

	// time.Sleep(time.Second * 1)

}

func handler(w http.ResponseWriter, req *http.Request) {
	precision := req.FormValue("precision")
	fmt.Fprintf(w, "Precision: "+precision)

}

// func setupRoutes(app *fiber.App) {

// 	app.Post("/api/v1/precision", func(c *fiber.Ctx) {
// 		fmt.Println(c.Body())

// 	})

// }
