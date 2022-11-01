//go:build unit

package csv

import (
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"testing"
)

func TestProcessor_ParseCities(t *testing.T) {
	// given
	underTest := NewProcessor()

	// when
	actual := underTest.ParseCities()

	// then
	assert.Equal(t, 7904, len(actual))
	assert.Contains(t, actual, service.CityCSV{
		Name:     "RIMINI",
		Province: "RN",
		Code:     "H294",
	})
}

func TestProcessor_GetCityCodesCache(t *testing.T) {
	// given
	underTest := NewProcessor()

	// when
	actual := underTest.GetCityCodesCache(nil)

	// then
	assert.Empty(t, actual)
}

func TestProcessor_GetCityPlacesCache(t *testing.T) {
	// given
	underTest := NewProcessor()

	// when
	actual := underTest.GetCityPlacesCache(nil)

	// then
	assert.Empty(t, actual)
}
