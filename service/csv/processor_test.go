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
	// given

	// when

	// then
}
