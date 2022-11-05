//go:build unit

package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"taxcode-converter/service/mocks"
	"testing"
)

type dependencies struct {
	mockProcessor *mocks.CsvProcessor
}

func setupTaxCodeService() (service.TaxCodeService, dependencies) {
	d := dependencies{mockProcessor: new(mocks.CsvProcessor)}
	underTest := NewTaxCodeService(CreateValidator(), d.mockProcessor)
	return underTest, d
}

func TestService_CalculatePersonDataSuccess(t *testing.T) {
	// given
	underTest, dependencies := setupTaxCodeService()
	input := service.CalculatePersonDataRequest{TaxCode: "BGNLSN93P19H294L"}
	dependencies.mockProcessor.On("CityFromCode", "H294").Return(&service.CityCSV{
		Name:     "RIMINI",
		Province: "RN",
		Code:     "H294",
	})

	// when
	actual, err := underTest.CalculatePersonData(input)

	// then
	assert.NotNil(t, actual)
	assert.Nil(t, err)
	expected := &service.CalculatePersonDataResponse{
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
	}
	assert.Equal(t, expected, actual)
	dependencies.mockProcessor.AssertExpectations(t)
}

func TestService_CalculatePersonDataReturnsErrorWhenNotValidRequest(t *testing.T) {
	// given
	underTest, dependencies := setupTaxCodeService()
	input := service.CalculatePersonDataRequest{TaxCode: "dummyTaxCode"}

	// when
	actual, err := underTest.CalculatePersonData(input)

	// then
	assert.Nil(t, actual)
	assert.NotNil(t, err)
	expected := service.NewValidationError("taxCode must be a valid tax code")
	assert.Equal(t, expected, err)
	dependencies.mockProcessor.AssertExpectations(t)
}

// TODO to fix all the tests
func TestService_CalculateTaxCode(t *testing.T) {
	assert.True(t, true)
}
