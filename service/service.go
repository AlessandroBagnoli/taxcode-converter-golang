package service

//go:generate mockery --all --output $PWD/mocks
import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CalculateTaxCode(c *fiber.Ctx) error
	CalculatePersonData(c *fiber.Ctx) error
	HandleError(c *fiber.Ctx, err error) error
}

type TaxCodeService interface {
	CalculateTaxCode(c context.Context, req CalculateTaxCodeRequest) (*CalculateTaxCodeResponse, error)
	CalculatePersonData(c context.Context, req CalculatePersonDataRequest) (*CalculatePersonDataResponse, error)
}

type Validator interface {
	ValidateCalculateTaxCodeReq(req CalculateTaxCodeRequest) error
	ValidateCalculatePersonDataReq(req CalculatePersonDataRequest) error
}
