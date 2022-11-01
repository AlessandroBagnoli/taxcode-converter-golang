//go:build unit

package csv

import (
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"testing"
)

func TestNewProcessor(t *testing.T) {
	// given

	// when
	actual := NewProcessor()

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
