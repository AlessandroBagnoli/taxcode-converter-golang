package validation

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"taxcode-converter/service"
	"time"
)

type Validator struct {
	validator validator.Validate
}

func NewValidator(v validator.Validate) Validator {
	return Validator{v}
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
		return false
	}
}

func ValidTaxCode(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		ok, err := regexp.MatchString("^([A-Z]{6}[0-9LMNPQRSTUV]{2}[ABCDEHLMPRST][0-9LMNPQRSTUV]{2}[A-Z][0-9LMNPQRSTUV]{3}[A-Z])$|(\\d{11})$", field.String())
		if err != nil {
			log.Warn(err)
		}
		return ok
	default:
		return false
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
