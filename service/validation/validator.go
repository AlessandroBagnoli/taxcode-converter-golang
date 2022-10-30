package validation

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"reflect"
	"taxcode-converter/service"
	"time"
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

func DateInThePast(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Type() {
	case reflect.TypeOf(time.Time{}):
		casted := civil.DateOf(field.Interface().(time.Time))
		return casted.Before(civil.DateOf(time.Now()))
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func TimeValue(v reflect.Value) interface{} {
	//TODO extend it for all the types inside civil package
	n, ok := v.Interface().(civil.Date)

	if !ok {
		return nil
	}

	return n.In(time.UTC)
}
