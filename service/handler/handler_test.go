//go:build functional

package handler

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandler_CalculateTaxCode(t *testing.T) {
	h := InitDependencies()
	app := CreateFiberApp(h)

	tests := []struct {
		description  string
		route        string
		body         string
		expectedCode int
		expectedBody string
	}{
		{
			description:  "should return 200 with calculated taxcode when request is valid",
			route:        "/api/v1/taxcode:calculate-tax-code",
			body:         `{"gender":"MALE","name":"Alessandro","surname":"Bagnoli","dateOfBirth":"1993-09-19","birthPlace":"Rimini","province":"RN"}`,
			expectedCode: 200,
			expectedBody: `{"taxCode":"BGNLSN93P19H294L"}`,
		},
		{
			description:  "should return 400 when body is not validated",
			route:        "/api/v1/taxcode:calculate-tax-code",
			body:         `{"gender":"MALE","name":"Alessandro","surname":"Bagnoli","dateOfBirth":"2022-10-29","birthPlace":"","province":"RN"}`,
			expectedCode: 400,
			expectedBody: `{"type":"about:blank","title":"Bad Request","status":400,"detail":"Key: 'CalculateTaxCodeRequest.BirthPlace' Error:Field validation for 'BirthPlace' failed on the 'required' tag","instance":"/api/v1/taxcode:calculate-tax-code"}`,
		},
		{
			description:  "should return 400 when body cannot be unmarshalled into go value",
			route:        "/api/v1/taxcode:calculate-tax-code",
			body:         `"something I cannot unmarshal into go value"`,
			expectedCode: 400,
			expectedBody: `{"type":"about:blank","title":"Bad Request","status":400,"detail":"json: cannot unmarshal string into Go value of type service.CalculateTaxCodeRequest","instance":"/api/v1/taxcode:calculate-tax-code"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// given
			req := httptest.NewRequest(fiber.MethodPost, test.route, bytes.NewReader([]byte(test.body)))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			// when
			resp, err := app.Test(req)

			// then
			assert.Nil(t, err)
			assert.Equal(t, test.expectedCode, resp.StatusCode)
			all, _ := io.ReadAll(resp.Body)
			actualBody := string(all)
			assert.Equal(t, test.expectedBody, actualBody)
		})
	}

	_ = app.Shutdown()

}
