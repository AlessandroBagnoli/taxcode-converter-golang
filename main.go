package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go-poc/service/handler"
	"go-poc/service/taxcode"
)

func main() {

	taxCodeService := taxcode.NewTaxCodeService()
	h := handler.NewHandler(taxCodeService)

	app := fiber.New()
	app.Post("/api/v1/taxcode:calculate-person-data", h.CalculatePersonData)

	log.Fatal(app.Listen(":8080"))
}
