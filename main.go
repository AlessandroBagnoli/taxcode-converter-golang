package main

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swaggerfiber "github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"strings"
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

	// creation of fiber app with custom config for error handling, usage of logger and cors
	app := fiber.New(fiber.Config{ErrorHandler: h.HandleError})
	app.Use(configureFiberLogger())
	app.Use(cors.New())

	// routing for swagger documentation
	app.Get("/swagger/*", swaggerfiber.HandlerDefault)

	// routing for apis
	v1 := app.Group("/api/v1")
	v1.Post("/taxcode:calculate-tax-code", h.CalculateTaxCode)
	v1.Post("/taxcode:calculate-person-data", h.CalculatePersonData)

	// server listening
	log.Fatal(app.Listen(":8080"))
}

func configureFiberLogger() fiber.Handler {
	return logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Path(), "/swagger/")
		},
		Format: "[${time}]|${status}|${resBody}|${latency} - ${method}|${path}|${body}\n",
	})
}

func configureValidator() validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("notblank", validators.NotBlank); err != nil {
		log.Fatal(err)
	}
	if err := validate.RegisterValidation("taxcode", validatorservice.IsValidTaxCode); err != nil {
		log.Fatal(err)
	}
	validate.RegisterCustomTypeFunc(validatorservice.TimeValue, civil.Date{}, civil.DateTime{}, civil.Time{})
	if err := validate.RegisterValidation("inthepast", validatorservice.DateInThePast); err != nil {
		log.Fatal(err)
	}
	return *validate
}
