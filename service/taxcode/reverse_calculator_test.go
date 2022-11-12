package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"testing"
)

func Test_reverseTaxCodeSuccess(t *testing.T) {
	tests := []struct {
		name         string
		inputTaxCode string
		cityCode     string
		expected     *service.CalculatePersonDataResponse
	}{
		{
			name:         "BGNLSN93P19H294L should return as expected",
			inputTaxCode: "BGNLSN93P19H294L",
			cityCode:     "H294",
			expected: &service.CalculatePersonDataResponse{
				Gender:  service.GenderMale,
				Name:    "LSN",
				Surname: "BGN",
				DateOfBirth: civil.Date{
					Year:  1993,
					Month: 9,
					Day:   19,
				},
				BirthPlace: "RIMINI",
				Province:   "RN",
				TaxCode:    "BGNLSN93P19H294L",
			},
		},
		{
			name:         "PTRRSL10R45G479I should return as expected",
			inputTaxCode: "PTRRSL10R45G479I",
			cityCode:     "G479",
			expected: &service.CalculatePersonDataResponse{
				Gender:  service.GenderFemale,
				Name:    "RSL",
				Surname: "PTR",
				DateOfBirth: civil.Date{
					Year:  2010,
					Month: 10,
					Day:   5,
				},
				BirthPlace: "PESARO",
				Province:   "PU",
				TaxCode:    "PTRRSL10R45G479I",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			extractor := func(s string) *service.CityCSV {
				return &service.CityCSV{
					Name:     tt.expected.BirthPlace,
					Province: tt.expected.Province,
					Code:     tt.cityCode,
				}
			}

			// when
			actual, err := reverseTaxCode(tt.inputTaxCode, extractor)

			// then
			assert.Nil(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_reverseTaxCodeShouldFailWhenNoCityFound(t *testing.T) {
	// given
	input := "BGNLSN93P19H294L"
	extractor := func(s string) *service.CityCSV { return nil }

	// when
	actual, err := reverseTaxCode(input, extractor)

	// then
	assert.Nil(t, actual)
	assert.NotNil(t, err)
	expected := service.NewCityNotPresentError("The city with code H294 does not exixts")
	assert.Equal(t, expected, err)
}
