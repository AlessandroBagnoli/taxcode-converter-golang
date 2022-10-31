package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"taxcode-converter/service"
)

type Service struct {
	validator validator.Validate
}

func NewTaxCodeService(v validator.Validate) Service {
	return Service{validator: v}
}

func (s Service) CalculateTaxCode(req service.CalculateTaxCodeRequest) (*service.CalculateTaxCodeResponse, error) {
	if err := ValidateReq(s.validator, req); err != nil {
		return nil, err
	}
	dummyResponse := &service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"}
	return dummyResponse, nil
}

func (s Service) CalculatePersonData(req service.CalculatePersonDataRequest) (*service.CalculatePersonDataResponse, error) {
	if err := ValidateReq(s.validator, req); err != nil {
		return nil, err
	}
	dummyResponse := &service.CalculatePersonDataResponse{
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
