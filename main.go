package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&Dto{
			Data:      "Hello world!",
			NestedDto: &NestedDto{NestedData: 56}})
	})

	log.Fatal(app.Listen(":8080"))
}
