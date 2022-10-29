package taxcode

import (
	"context"
	"go-poc/service"
	"time"
)

type Service struct {
}

func NewTaxCodeService() *Service {
	return &Service{}
}

func (s *Service) CalculateTaxCode(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CalculatePersonData(ctx context.Context, req *service.CalculatePersonDataRequest) (*service.CalculatePersonDataResponse, error) {
	dummyResponse := &service.CalculatePersonDataResponse{
		Gender:      service.GenderMale,
		Name:        "Alessandro",
		Surname:     "Bagnoli",
		DateOfBirth: time.Time{},
		BirthPlace:  "Rimini",
		Province:    "RN",
		TaxCode:     "BGNLSN93P19H294L",
	}
	return dummyResponse, nil
}
