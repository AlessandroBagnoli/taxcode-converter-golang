package validator

import (
	"github.com/go-playground/validator/v10"
	"taxcode-converter/service"
)

type Validator struct {
	validator validator.Validate
}

func NewValidator(v validator.Validate) *Validator {
	return &Validator{v}
}

func (v Validator) ValidateCalculateTaxCodeReq(req service.CalculateTaxCodeRequest) error {
	return v.validator.Struct(req)
}

func (v Validator) ValidateCalculatePersonDataReq(req service.CalculatePersonDataRequest) error {
	return v.validator.Struct(req)
}
