package taxcode

import (
	"github.com/go-playground/validator/v10"
	"taxcode-converter/service"
)

type Service struct {
	validator validator.Validate
	processor service.CsvProcessor
}

func NewTaxCodeService(v validator.Validate, p service.CsvProcessor) Service {
	return Service{validator: v, processor: p}
}

func (s Service) CalculateTaxCode(req service.CalculateTaxCodeRequest) (*service.CalculateTaxCodeResponse, error) {
	if err := ValidateReq(s.validator, req); err != nil {
		return nil, err
	}
	return calculate(req, s.processor.CityFromPlace)
}

func (s Service) CalculatePersonData(req service.CalculatePersonDataRequest) (*service.CalculatePersonDataResponse, error) {
	if err := ValidateReq(s.validator, req); err != nil {
		return nil, err
	}
	return reverseTaxCode(req.TaxCode, s.processor.CityFromCode)
}
