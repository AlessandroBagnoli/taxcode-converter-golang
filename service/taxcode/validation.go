package taxcode

import (
	"cloud.google.com/go/civil"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
	"taxcode-converter/service"
	"time"
)

var tagMessagesMap = map[string]string{
	"required":  "%s is required",
	"notblank":  "%s must not be blank",
	"oneof":     "%s must be one of admitted values",
	"inthepast": "%s must be in the past",
	"taxcode":   "%s must be a valid tax code",
}

func ValidateReq[T service.GenericRequest](v validator.Validate, req T) error {
	var errs []string
	err := v.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg, ok := tagMessagesMap[err.Tag()]
			if !ok {
				msg = err.Tag()
			}
			errorMsg := fmt.Sprintf(msg, err.Field())
			errs = append(errs, errorMsg)
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
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
