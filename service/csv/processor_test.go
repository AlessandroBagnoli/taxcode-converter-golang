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
