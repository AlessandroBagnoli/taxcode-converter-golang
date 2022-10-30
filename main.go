package main

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/gofiber/fiber/v2"
	swaggerfiber "github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	_ "taxcode-converter/docs"
	"taxcode-converter/service/handler"
	"taxcode-converter/service/taxcode"
	validatorservice "taxcode-converter/service/validation"
)

// @title taxcode-converter
// @version 1.0
// @host localhost:8080
// @BasePath /
//
//go:generate swag init --pd
func main() {

	// dependencies injection
	taxCodeService := taxcode.NewTaxCodeService()
	validate := configureValidator()
	v := validatorservice.NewValidator(validate)
	h := handler.NewHandler(taxCodeService, v)

	// creation of fiber app with custom config for error handling
	app := fiber.New(fiber.Config{ErrorHandler: h.HandleError})

	// routing for swagger documentation
	app.Get("/swagger/*", swaggerfiber.HandlerDefault)

	// routing for apis
	v1 := app.Group("/api/v1")
	v1.Post("/taxcode:calculate-tax-code", h.CalculateTaxCode)
	v1.Post("/taxcode:calculate-person-data", h.CalculatePersonData)

	// server listening
	log.Fatal(app.Listen(":8080"))
}

func configureValidator() validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		log.Fatal(err)
	}
	validate.RegisterCustomTypeFunc(validatorservice.TimeValue, civil.Date{}, civil.DateTime{}, civil.Time{})
	err = validate.RegisterValidation("inthepast", validatorservice.DateInThePast)
	if err != nil {
		log.Fatal(err)
	}
	return *validate
}
