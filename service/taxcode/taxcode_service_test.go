//go:build unit

package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"taxcode-converter/service"
	"taxcode-converter/service/mocks"
	"testing"
)

type TaxCodeServiceTestSuite struct {
	suite.Suite
	validate  validator.Validate
	processor *mocks.CsvProcessor
	underTest service.TaxCodeService
}

func (suite *TaxCodeServiceTestSuite) SetupTest() {
	suite.validate = CreateValidator()
	suite.processor = new(mocks.CsvProcessor)
	suite.underTest = NewTaxCodeService(suite.validate, suite.processor)
}

func (suite *TaxCodeServiceTestSuite) TearDownTest() {
	suite.processor.AssertExpectations(suite.T())
}

func (suite *TaxCodeServiceTestSuite) Test_CalculatePersonDataSuccess() {
	input := service.CalculatePersonDataRequest{TaxCode: "BGNLSN93P19H294L"}
	suite.processor.On("CityFromCode", "H294").Return(&service.CityCSV{
		Name:     "RIMINI",
		Province: "RN",
		Code:     "H294",
	})

	// when
	actual, err := suite.underTest.CalculatePersonData(input)

	// then
	suite.NotNil(actual)
	suite.Nil(err)
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
	suite.Equal(expected, actual)
}

func (suite *TaxCodeServiceTestSuite) Test_CalculatePersonDataReturnsErrorWhenNotValidRequest() {
	// given
	input := service.CalculatePersonDataRequest{TaxCode: "dummyTaxCode"}

	// when
	actual, err := suite.underTest.CalculatePersonData(input)

	// then
	suite.Nil(actual)
	suite.NotNil(err)
	expected := service.NewValidationError("taxCode must be a valid tax code")
	suite.Equal(expected, err)
}

func (suite *TaxCodeServiceTestSuite) Test_CalculateTaxCodeSuccess() {
	// given
	input := service.CalculateTaxCodeRequest{
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
	}
	suite.processor.On("CityFromPlace", service.Place{CityName: "RIMINI", Province: "RN"}).
		Return(&service.CityCSV{Name: "RIMINI", Province: "RN", Code: "H294"})

	// when
	actual, err := suite.underTest.CalculateTaxCode(input)

	// then
	suite.Nil(err)
	suite.NotNil(actual)
	expected := &service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"}
	suite.Equal(expected, actual)
}

func (suite *TaxCodeServiceTestSuite) Test_CalculateTaxCodeReturnsErrorWhenCalculateFailsDueToCityNotPresent() {
	// given
	input := service.CalculateTaxCodeRequest{
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
	}
	suite.processor.On("CityFromPlace", service.Place{CityName: "RIMINI", Province: "RN"}).Return(nil)

	// when
	actual, err := suite.underTest.CalculateTaxCode(input)

	// then
	suite.Nil(actual)
	suite.NotNil(err)
	expected := service.NewCityNotPresentError("The city Rimini and province RN do not exixt")
	suite.Equal(expected, err)
}

func (suite *TaxCodeServiceTestSuite) Test_CalculateTaxCodeReturnsErrorWhenNoValidRequest() {
	// given
	input := service.CalculateTaxCodeRequest{
		Gender:  service.GenderMale,
		Name:    "",
		Surname: "",
		DateOfBirth: civil.Date{
			Year:  1993,
			Month: 9,
			Day:   19,
		},
		BirthPlace: "Rimini",
		Province:   "RN",
	}

	// when
	actual, err := suite.underTest.CalculateTaxCode(input)

	// then
	suite.Nil(actual)
	suite.NotNil(err)
	expected := service.NewValidationError("name is required, surname is required")
	suite.Equal(expected, err)
}

func Test_TaxCodeService(t *testing.T) {
	suite.Run(t, new(TaxCodeServiceTestSuite))
}
