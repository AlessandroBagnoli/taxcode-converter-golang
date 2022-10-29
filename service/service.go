package service

//go:generate mockery --all --output $PWD/mocks
import "context"

type TaxCodeService interface {
	CalculateTaxCode(ctx context.Context)
	CalculatePersonData(ctx context.Context, req *CalculatePersonDataRequest) (*CalculatePersonDataResponse, error)
}
