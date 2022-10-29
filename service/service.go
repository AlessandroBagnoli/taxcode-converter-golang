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
	CalculateTaxCode(ctx context.Context, req *CalculateTaxCodeRequest) (*CalculateTaxCodeResponse, error)
	CalculatePersonData(ctx context.Context, req *CalculatePersonDataRequest) (*CalculatePersonDataResponse, error)
}
