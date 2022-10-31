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

func TestValidateCalculateTaxCodeRequest(t *testing.T) {
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
		{
			description: "should return error when CalculateTaxCodeRequest is not ok",
			input: service.CalculateTaxCodeRequest{
				Gender:  "MALE",
				Name:    " ",
				Surname: "Rossi",
				DateOfBirth: civil.Date{
					Year:  2000,
					Month: 1,
					Day:   1,
				},
				BirthPlace: "Roma",
				Province:   "RM",
			},
			expected: service.NewRuntimeError(400, "name must not be blank"),
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

func TestValidateCalculatePersonDataRequest(t *testing.T) {
	v := setup()

	tests := []struct {
		description string
		input       service.CalculatePersonDataRequest
		expected    error
	}{
		{
			description: "should return nil when CalculatePersonDataRequest is ok",
			input:       service.CalculatePersonDataRequest{TaxCode: "TSMFZF92H15F081V"},
			expected:    nil,
		},
		{
			description: "should return error when CalculateTaxCodeRequest is not ok",
			input:       service.CalculatePersonDataRequest{TaxCode: "dummyTaxCode"},
			expected:    service.NewRuntimeError(400, "taxCode must be a valid tax code"),
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
