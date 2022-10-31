//go:build unit

package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"testing"
)

func setup() validator.Validate {
	return CreateValidator()
}

func TestValidateReq(t *testing.T) {
	v := setup()

	tests := []struct {
		description string
		input       service.CalculateTaxCodeRequest
		expected    error
	}{
		{
			description: "should return nil when CalculateTaxCodeRequest is ok",
			input: service.CalculateTaxCodeRequest{
				Gender:  "MALE",
				Name:    "Mario",
				Surname: "Rossi",
				DateOfBirth: civil.Date{
					Year:  2000,
					Month: 1,
					Day:   1,
				},
				BirthPlace: "Roma",
				Province:   "RM",
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// given

			// when
			actual := ValidateReq(v, test.input)

			// then
			assert.Equal(t, test.expected, actual)
		})
	}
}
