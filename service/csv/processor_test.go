//go:build unit

package csv

import (
	"github.com/stretchr/testify/assert"
	"os"
	"taxcode-converter/service"
	"testing"
)

func TestNewProcessor(t *testing.T) {
	// given
	file, _ := os.ReadFile("../../assets/italian-cities.csv")

	// when
	actual := NewProcessor(file)

	// then
	assert.Equal(t, 7904, len(actual.cities))
	assert.Contains(t, actual.cities, service.CityCSV{
		Name:     "RIMINI",
		Province: "RN",
		Code:     "H294",
	})
	assert.NotEmpty(t, actual.cityPlacesCache)
	assert.NotEmpty(t, actual.cityCodesCache)
}
