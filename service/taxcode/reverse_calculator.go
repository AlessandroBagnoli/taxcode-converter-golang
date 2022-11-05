package taxcode

import (
	"cloud.google.com/go/civil"
	"fmt"
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

func reverseTaxCode(tc string, cityExtractor func(string) *service.CityCSV) (*service.CalculatePersonDataResponse, error) {
	taxCode := strings.ToUpper(tc)
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
	month := charMonthMap[m]

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
	city := cityExtractor(cityCode)
	if city == nil {
		return nil, service.NewCityNotPresentError(fmt.Sprintf("The city with code %s does not exixts", cityCode))
	}

	personData := &service.CalculatePersonDataResponse{
		Gender:      service.Gender(gender),
		Name:        name,
		Surname:     surname,
		DateOfBirth: birthDate,
		BirthPlace:  strings.ToUpper(city.Name),
		Province:    strings.ToUpper(city.Province),
		TaxCode:     taxCode,
	}
	return personData, nil
}
