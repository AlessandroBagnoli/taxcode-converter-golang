package service

//go:generate mockery --all --output $PWD/mocks
import (
	"github.com/gofiber/fiber/v2"
)

// Handler declares all the handlers methods used to process http incoming requests
type Handler interface {
	CalculateTaxCode(c *fiber.Ctx) error
	CalculatePersonData(c *fiber.Ctx) error
	HandleError(c *fiber.Ctx, err error) error
}

// TaxCodeService declares the business logic methods
type TaxCodeService interface {
	CalculateTaxCode(req CalculateTaxCodeRequest) (*CalculateTaxCodeResponse, error)
	CalculatePersonData(req CalculatePersonDataRequest) (*CalculatePersonDataResponse, error)
}

// CsvProcessor declares utility methods to get data from the parsed csv of cities
type CsvProcessor interface {
	CityFromCode(code string) *CityCSV
	CityFromPlace(place Place) *CityCSV
}
