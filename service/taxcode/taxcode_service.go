package taxcode

import (
	"cloud.google.com/go/civil"
	"context"
	"taxcode-converter/service"
)

type Service struct {
}

func NewTaxCodeService() *Service {
	return &Service{}
}

func (s Service) CalculateTaxCode(c context.Context, req service.CalculateTaxCodeRequest) (service.CalculateTaxCodeResponse, error) {
	dummyResponse := service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"}
	return dummyResponse, nil
}

func (s Service) CalculatePersonData(c context.Context, req service.CalculatePersonDataRequest) (service.CalculatePersonDataResponse, error) {
	dummyResponse := service.CalculatePersonDataResponse{
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
		TaxCode:    "BGNLSN93P19H294L",
	}
	return dummyResponse, nil
}
