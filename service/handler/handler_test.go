package handler

import (
	"bytes"
	"cloud.google.com/go/civil"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"taxcode-converter/service"
	"testing"
)

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
			h := InitDependencies()
			app := CreateFiberApp(h)
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
