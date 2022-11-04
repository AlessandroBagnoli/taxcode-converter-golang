package taxcode

import (
	"cloud.google.com/go/civil"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
	"taxcode-converter/service"
	"time"
)

var charMonthMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"H": 6,
	"L": 7,
	"M": 8,
	"P": 9,
	"R": 10,
	"S": 11,
	"T": 12,
}

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
	// TODO implement BL
	dummyResponse := &service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"}
	return dummyResponse, nil
}

func (s Service) CalculatePersonData(req service.CalculatePersonDataRequest) (*service.CalculatePersonDataResponse, error) {
	if err := ValidateReq(s.validator, req); err != nil {
		return nil, err
	}
	trueResponse := reverseTaxCode(req, s)
	return trueResponse, nil
}

// TODO to extract?
func reverseTaxCode(req service.CalculatePersonDataRequest, s Service) *service.CalculatePersonDataResponse {
	taxCode := strings.ToUpper(req.TaxCode)
	// surname + name
	surname := taxCode[0:3]
	name := taxCode[3:6]

	// day + gender
	sDay := taxCode[9:11]
	day, _ := strconv.Atoi(sDay)
	gender := service.GenderMale
	dayToConsider := day
	if day > 31 {
		gender = service.GenderFemale
		dayToConsider = day - 40
	}

	// month
	m := taxCode[8:9]
	month, ok := charMonthMap[m]
	if !ok {
		month = 0
	}

	// year
	lastTwoDigitsThisYear := civil.DateOf(time.Now()).Year % 1e2
	yy := taxCode[6:8]
	y, _ := strconv.Atoi(yy)
	theYear := 2000 + y
	if y > lastTwoDigitsThisYear {
		theYear = 1900 + y
	}

	birthDate := civil.Date{
		Year:  theYear,
		Month: time.Month(month),
		Day:   dayToConsider,
	}

	// city
	cityCode := taxCode[11:15]
	city := s.processor.CityFromCode(cityCode)

	trueResponse := &service.CalculatePersonDataResponse{
		Gender:      service.Gender(gender),
		Name:        name,
		Surname:     surname,
		DateOfBirth: birthDate,
		BirthPlace:  strings.ToUpper(city.Name),
		Province:    strings.ToUpper(city.Province),
		TaxCode:     taxCode,
	}
	return trueResponse
}
