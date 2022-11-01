//go:build unit

package csv

import (
	"github.com/stretchr/testify/assert"
	"os"
	"taxcode-converter/service"
	"testing"
)

func setup() Processor {
	file, _ := os.ReadFile("../../assets/italian-cities.csv")
	return NewProcessor(file)
}

func TestNewProcessor(t *testing.T) {
	// given
	actual := setup()

	// when

	// then
	assert.Equal(t, 7904, len(actual.cities))
	assert.Contains(t, actual.cities, &service.CityCSV{
		Name:     "RIMINI",
		Province: "RN",
		Code:     "H294",
	})
	assert.NotEmpty(t, actual.cityPlacesCache)
	assert.NotEmpty(t, actual.cityCodesCache)
}

func TestProcessor_CityFromCode(t *testing.T) {
	processor := setup()

	tests := []struct {
		name     string
		input    string
		expected *service.CityCSV
	}{
		{name: "H294 should return RIMINI", input: "H294", expected: &service.CityCSV{Name: "RIMINI", Province: "RN", Code: "H294"}},
		{name: "A074 should return AGLIÈ", input: "A074", expected: &service.CityCSV{Name: "AGLIÈ", Province: "TO", Code: "A074"}},
		{name: "A275 should return ANDEZENO", input: "A275", expected: &service.CityCSV{Name: "ANDEZENO", Province: "TO", Code: "A275"}},
		{name: "M333 should return SANT'OMOBONO TERME", input: "M333", expected: &service.CityCSV{Name: "SANT'OMOBONO TERME", Province: "BG", Code: "M333"}},
		{name: "code not existing should return nil", input: "dummyCode", expected: nil},
		{name: "empty string should return nil", input: "", expected: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// given

			// when
			actual := processor.CityFromCode(test.input)

			// then
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestProcessor_CityFromPlace(t *testing.T) {
	processor := setup()

	tests := []struct {
		name     string
		input    service.Place
		expected *service.CityCSV
	}{
		{name: "RIMINI should return RIMINI", input: service.Place{CityName: "RIMINI", Province: "RN"}, expected: &service.CityCSV{Name: "RIMINI", Province: "RN", Code: "H294"}},
		{name: "AGLIÈ should return AGLIÈ", input: service.Place{CityName: "AGLIÈ", Province: "TO"}, expected: &service.CityCSV{Name: "AGLIÈ", Province: "TO", Code: "A074"}},
		{name: "ANDEZENO should return ANDEZENO", input: service.Place{CityName: "ANDEZENO", Province: "TO"}, expected: &service.CityCSV{Name: "ANDEZENO", Province: "TO", Code: "A275"}},
		{name: "SANT'OMOBONO TERME should return SANT'OMOBONO TERME", input: service.Place{CityName: "SANT'OMOBONO TERME", Province: "BG"}, expected: &service.CityCSV{Name: "SANT'OMOBONO TERME", Province: "BG", Code: "M333"}},
		{name: "empty city and province should return nil", input: service.Place{CityName: "", Province: ""}, expected: nil},
		{name: "empty city should return nil", input: service.Place{CityName: "", Province: "RN"}, expected: nil},
		{name: "empty province should return nil", input: service.Place{CityName: "RIMINI", Province: ""}, expected: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// given

			// when
			actual := processor.CityFromPlace(test.input)

			// then
			assert.Equal(t, test.expected, actual)
		})
	}
}
