package handler

import (
	"bytes"
	"cloud.google.com/go/civil"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swaggerfiber "github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"strings"
	"taxcode-converter/service"
	"taxcode-converter/service/taxcode"
	validatorservice "taxcode-converter/service/validation"
	"testing"
)

// TODO all these setup methods are just copied from main.go. Find a place where to put them in order to reuse them also in test without copying everytime
func setup() Handler {
	taxCodeService := taxcode.NewTaxCodeService()
	validate := configureValidator()
	v := validatorservice.NewValidator(validate)
	h := NewHandler(taxCodeService, v)
	return h
}

func configureValidator() validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("notblank", validators.NotBlank); err != nil {
		log.Fatal(err)
	}
	if err := validate.RegisterValidation("taxcode", validatorservice.ValidTaxCode); err != nil {
		log.Fatal(err)
	}
	validate.RegisterCustomTypeFunc(validatorservice.TimeValue, civil.Date{}, civil.DateTime{}, civil.Time{})
	if err := validate.RegisterValidation("inthepast", validatorservice.DateInThePast); err != nil {
		log.Fatal(err)
	}
	return *validate
}

func createFiberApp(h Handler) *fiber.App {
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

func TestHandler_CalculateTaxCode(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		body         service.CalculateTaxCodeRequest
		expectedCode int
		expectedBody service.CalculateTaxCodeResponse
	}{
		{
			description: "should return 200 with body as expected",
			route:       "/api/v1/taxcode:calculate-tax-code",
			body: service.CalculateTaxCodeRequest{
				Gender:  "MALE",
				Name:    "Alessandro",
				Surname: "Bagnoli",
				DateOfBirth: civil.Date{
					Year:  1993,
					Month: 9,
					Day:   19,
				},
				BirthPlace: "Rimini",
				Province:   "RN",
			},
			expectedCode: 200,
			expectedBody: service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"},
		},
		{
			description: "should return 200 with body as expected",
			route:       "/api/v1/taxcode:calculate-tax-code",
			body: service.CalculateTaxCodeRequest{
				Gender:  "MALE",
				Name:    "Alessandro",
				Surname: "Bagnoli",
				DateOfBirth: civil.Date{
					Year:  1993,
					Month: 9,
					Day:   19,
				},
				BirthPlace: "Rimini",
				Province:   "RN",
			},
			expectedCode: 200,
			expectedBody: service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// given
			h := setup()
			app := createFiberApp(h)
			marshal, _ := json.Marshal(test.body)
			req := httptest.NewRequest(fiber.MethodPost, test.route, bytes.NewReader(marshal))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			// when
			resp, err := app.Test(req)

			// then
			assert.Nil(t, err)
			assert.Equal(t, test.expectedCode, resp.StatusCode)
			all, err := io.ReadAll(resp.Body)
			actualBody := new(service.CalculateTaxCodeResponse)
			if err = json.Unmarshal(all, actualBody); err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, test.expectedBody, *actualBody)
		})
	}

}
