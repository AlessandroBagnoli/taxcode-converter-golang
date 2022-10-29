package service

//go:generate mockery --all --output $PWD/mocks
import "context"

type TaxCodeService interface {
	CalculateTaxCode(ctx context.Context, req *CalculateTaxCodeRequest) (*CalculateTaxCodeResponse, error)
	CalculatePersonData(ctx context.Context, req *CalculatePersonDataRequest) (*CalculatePersonDataResponse, error)
}
