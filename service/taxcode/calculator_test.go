//go:build unit

package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"testing"
)

func TestService_calculateTaxCodeSuccess(t *testing.T) {
	tests := []struct {
		name     string
		input    service.CalculateTaxCodeRequest
		cityCode string
		expected string
	}{
		{"testcase 1", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "Alessandro",
			Surname: "Bagnoli",
			DateOfBirth: civil.Date{
				Year:  1993,
				Month: 9,
				Day:   19,
			},
			BirthPlace: "Rimini",
			Province:   "RN",
		}, "H294", "BGNLSN93P19H294L"},
		{"testcase 2", service.CalculateTaxCodeRequest{
			Gender:  service.GenderFemale,
			Name:    "F",
			Surname: "F",
			DateOfBirth: civil.Date{
				Year:  1993,
				Month: 9,
				Day:   19,
			},
			BirthPlace: "Rimini",
			Province:   "RN",
		}, "H294", "FXXFXX93P59H294P"},
		{"testcase 3", service.CalculateTaxCodeRequest{
			Gender:  service.GenderFemale,
			Name:    "A",
			Surname: "FF",
			DateOfBirth: civil.Date{
				Year:  1993,
				Month: 9,
				Day:   19,
			},
			BirthPlace: "Rimini",
			Province:   "RN",
		}, "H294", "FFXAXX93P59H294S"},
		{"testcase 4", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "A",
			Surname: "AIO",
			DateOfBirth: civil.Date{
				Year:  2010,
				Month: 10,
				Day:   19,
			},
			BirthPlace: "roma",
			Province:   "rm",
		}, "H501", "AIOAXX10R19H501O"},
		{"testcase 5", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "AA",
			Surname: "",
			DateOfBirth: civil.Date{
				Year:  2010,
				Month: 10,
				Day:   19,
			},
			BirthPlace: "roma",
			Province:   "rm",
		}, "H501", "XXXAAX10R19H501R"},
		{"testcase 6", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "bae",
			Surname: "ba",
			DateOfBirth: civil.Date{
				Year:  1950,
				Month: 10,
				Day:   19,
			},
			BirthPlace: "pesaro",
			Province:   "pu",
		}, "G479", "BAXBAE50R19G479N"},
		{"testcase 7", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "ba",
			Surname: "baed",
			DateOfBirth: civil.Date{
				Year:  1950,
				Month: 10,
				Day:   19,
			},
			BirthPlace: "pesaro",
			Province:   "pu",
		}, "G479", "BDABAX50R19G479L"},
		{"testcase 8", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "qqq",
			Surname: "rossi",
			DateOfBirth: civil.Date{
				Year:  1950,
				Month: 10,
				Day:   5,
			},
			BirthPlace: "pesaro",
			Province:   "pu",
		}, "G479", "RSSQQQ50R05G479X"},
		{"testcase 9", service.CalculateTaxCodeRequest{
			Gender:  service.GenderMale,
			Name:    "hfg4",
			Surname: "5555",
			DateOfBirth: civil.Date{
				Year:  1993,
				Month: 9,
				Day:   19,
			},
			BirthPlace: "miagliano",
			Province:   "bi",
		}, "F189", "XXXHFG93P19F189A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			extractor := func(s service.Place) *service.CityCSV {
				return &service.CityCSV{
					Name:     "does not matter",
					Province: "does not matter",
					Code:     tt.cityCode,
				}
			}

			// when
			actual, err := calculate(tt.input, extractor)

			// then
			assert.Nil(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestService_CalculateTaxCode(t *testing.T) {
	// given
	input := service.CalculateTaxCodeRequest{
		Gender:  service.GenderMale,
		Name:    "Someone",
		Surname: "Someone",
		DateOfBirth: civil.Date{
			Year:  2000,
			Month: 10,
			Day:   10,
		},
		BirthPlace: "some place",
		Province:   "some province",
	}
	extractor := func(place service.Place) *service.CityCSV { return nil }

	// when
	actual, err := calculate(input, extractor)

	// then
	assert.Empty(t, actual)
	assert.NotNil(t, err)
	expected := service.NewCityNotPresentError("The city some place and province some province do not exixt")
	assert.Equal(t, expected, err)
}
