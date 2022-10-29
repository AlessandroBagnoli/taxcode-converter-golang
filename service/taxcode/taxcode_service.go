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

func (s *Service) CalculateTaxCode(ctx context.Context, req *service.CalculateTaxCodeRequest) (*service.CalculateTaxCodeResponse, error) {
	dummyResponse := &service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"}
	return dummyResponse, nil
}

func (s *Service) CalculatePersonData(ctx context.Context, req *service.CalculatePersonDataRequest) (*service.CalculatePersonDataResponse, error) {
	dummyResponse := &service.CalculatePersonDataResponse{
		Gender:      service.GenderMale,
		Name:        "Alessandro",
		Surname:     "Bagnoli",
		DateOfBirth: service.CivilTime(time.Date(1993, 9, 19, 0, 0, 0, 0, time.UTC)),
		BirthPlace:  "Rimini",
		Province:    "RN",
		TaxCode:     "BGNLSN93P19H294L",
	}
	return dummyResponse, nil
}
