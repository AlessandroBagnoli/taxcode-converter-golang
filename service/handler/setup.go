package handler

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swaggerfiber "github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
	"taxcode-converter/service/taxcode"
	"time"
)

// InitDependencies creates and injects dependencies, returns the handler for incoming http requests
func InitDependencies() Handler {
	v := createValidator()
	t := taxcode.NewTaxCodeService(v)
	h := NewHandler(t)
	return h
}

// CreateFiberApp creates and configure the fiber server
func CreateFiberApp(h Handler) *fiber.App {
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
	return app
}

func configureFiberLogger() fiber.Handler {
	return logger.New(logger.Config{
		// this is because I just want to log incoming request to the exposed api and ignore everything else
		Next: func(c *fiber.Ctx) bool {
			return !strings.Contains(c.Path(), "/api/v1/")
		},
		Format: "[${time}]|${status}|${resBody}|${latency} - ${method}|${path}|${body}\n",
	})
}

func createValidator() validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("notblank", validators.NotBlank); err != nil {
		log.Fatal(err)
	}
	if err := validate.RegisterValidation("taxcode", validTaxCode); err != nil {
		log.Fatal(err)
	}
	validate.RegisterCustomTypeFunc(timeValue, civil.Date{}, civil.DateTime{})
	if err := validate.RegisterValidation("inthepast", dateInThePast); err != nil {
		log.Fatal(err)
	}
	//to use the names which have been specified for JSON representations of structs, rather than normal Go field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	return *validate
}

func dateInThePast(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Type() {
	case reflect.TypeOf(time.Time{}):
		casted := civil.DateOf(field.Interface().(time.Time))
		return casted.Before(civil.DateOf(time.Now()))
	default:
		return false
	}
}

func validTaxCode(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		ok, err := regexp.MatchString("^([A-Z]{6}[0-9LMNPQRSTUV]{2}[ABCDEHLMPRST][0-9LMNPQRSTUV]{2}[A-Z][0-9LMNPQRSTUV]{3}[A-Z])$|(\\d{11})$", field.String())
		if err != nil {
			log.Warn(err)
		}
		return ok
	default:
		return false
	}
}

func timeValue(v reflect.Value) interface{} {
	switch v.Interface().(type) {
	case civil.Date:
		date := v.Interface().(civil.Date)
		return date.In(time.UTC)
	case civil.DateTime:
		dateTime := v.Interface().(civil.DateTime)
		return dateTime.In(time.UTC)
	default:
		return nil
	}

}
