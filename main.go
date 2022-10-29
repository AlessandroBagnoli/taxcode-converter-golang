package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	swaggerfiber "github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	_ "taxcode-converter/docs"
	"taxcode-converter/service/handler"
	"taxcode-converter/service/taxcode"
	validatorservice "taxcode-converter/service/validator"
)

// @title taxcode-converter
// @version 1.0
// @host localhost:8080
// @BasePath /
//
//go:generate swag init --pd
func main() {

	taxCodeService := taxcode.NewTaxCodeService()
	v := validatorservice.NewValidator(*validator.New())
	h := handler.NewHandler(taxCodeService, v)

	app := fiber.New(fiber.Config{ErrorHandler: h.HandleError})
	swagger := app.Group("/swagger")
	swagger.Get("/*", swaggerfiber.HandlerDefault)

	v1 := app.Group("/api/v1")
	v1.Post("/taxcode:calculate-tax-code", h.CalculateTaxCode)
	v1.Post("/taxcode:calculate-person-data", h.CalculatePersonData)

	log.Fatal(app.Listen(":8080"))
}
