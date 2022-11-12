package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func CreateValidator() validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("notblank", validators.NotBlank); err != nil {
		log.Fatal(err)
	}
	if err := validate.RegisterValidation("taxcode", validTaxCode); err != nil {
		log.Fatal(err)
	}
	validate.RegisterCustomTypeFunc(timeValue, civil.Date{}, civil.DateTime{})
	if err := validate.RegisterValidation("inthepast", dateInThePast); err != nil {
		log.Fatal(err)
	}
	//to use the names which have been specified for JSON representations of structs, rather than normal Go field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	return *validate
}

func dateInThePast(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Type() {
	case reflect.TypeOf(time.Time{}):
		casted := civil.DateOf(field.Interface().(time.Time))
		return casted.Before(civil.DateOf(time.Now()))
	default:
		return false
	}
}

func validTaxCode(fl validator.FieldLevel) bool {
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

func timeValue(v reflect.Value) interface{} {
	switch v.Interface().(type) {
	case civil.Date:
		date := v.Interface().(civil.Date)
		return date.In(time.UTC)
	case civil.DateTime:
		dateTime := v.Interface().(civil.DateTime)
		return dateTime.In(time.UTC)
	default:
		return nil
	}

}
