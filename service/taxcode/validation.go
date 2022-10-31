package taxcode

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"taxcode-converter/service"
)

var tagMessagesMap = map[string]string{
	"required":  "%s is required",
	"notblank":  "%s must not be blank",
	"oneof":     "%s must be one of admitted values",
	"inthepast": "%s must be in the past",
	"taxcode":   "%s must be a valid tax code",
}

func ValidateReq[T GenericRequest](v validator.Validate, req T) error {
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

type GenericRequest interface {
	service.CalculateTaxCodeRequest | service.CalculatePersonDataRequest
}
