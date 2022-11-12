package taxcode

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"taxcode-converter/service"
	"time"
)

func calculate(req service.CalculateTaxCodeRequest, cityExtractor func(place service.Place) *service.CityCSV) (string, error) {
	var fiscalCode bytes.Buffer
	regex := regexp.MustCompile("[^A-Z]")
	fcSurname := regex.ReplaceAllString(strings.ToUpper(req.Surname), "")
	fcName := regex.ReplaceAllString(strings.ToUpper(req.Name), "")
	fcBirthDate := req.DateOfBirth.In(time.UTC).Format("02-01-2006")

	// surname
	consonantsSurname := consonants(fcSurname)
	vowelsSurname := vowels(fcSurname)
	consonantsSurnameLength := len(consonantsSurname)
	fn, ok := surnameFunctionMap[consonantsSurnameLength]
	if !ok {
		fn = surnameCaseDefault
	}
	surname := fn(vowelsSurname, consonantsSurname)
	fiscalCode.WriteString(surname)

	// name
	consonantsName := consonants(fcName)
	vowelsName := vowels(fcName)
	consonantsNameLength := len(consonantsName)
	fn, ok = nameFunctionMap[consonantsNameLength]
	if !ok {
		fn = nameCaseDefault
	}
	name := fn(vowelsName, consonantsName)
	fiscalCode.WriteString(name)

	// year
	fiscalCode.WriteString(fcBirthDate[8:10])

	// month
	var month int
	if fcBirthDate[3:4] == "0" {
		month, _ = strconv.Atoi(fcBirthDate[4:5])
	} else {
		month, _ = strconv.Atoi(fcBirthDate[3:5])
	}
	fiscalCode.WriteString(monthCharMap[month])

	// day
	day, _ := strconv.Atoi(fcBirthDate[0:2])
	if req.Gender == service.GenderMale {
		if day < 10 {
			fiscalCode.WriteString("0" + strconv.FormatInt(int64(day), 10))
		} else {
			fiscalCode.WriteString(strconv.FormatInt(int64(day), 10))
		}
	} else {
		day += 40
		fiscalCode.WriteString(strconv.FormatInt(int64(day), 10))
	}

	// birth city
	cityName := req.BirthPlace
	province := req.Province
	place := service.Place{
		CityName: strings.ToUpper(cityName),
		Province: strings.ToUpper(province),
	}
	city := cityExtractor(place)
	if city == nil {
		return "", service.NewCityNotPresentError(fmt.Sprintf("The city %s and province %s do not exixt", cityName, province))
	}
	fiscalCode.WriteString(city.Code)

	// control char
	evenSum := reduce([]int{1, 3, 5, 7, 9, 11, 13}, 0, func(a int, b int) int { return a + evenSumMap[fiscalCode.String()[b:b+1]] })
	oddSum := reduce([]int{0, 2, 4, 6, 8, 10, 12, 14}, 0, func(a int, b int) int { return a + oddSumMap[fiscalCode.String()[b:b+1]] })
	controlInteger := (evenSum + oddSum) % 26
	controlCharacter := controlCharMap[controlInteger]
	fiscalCode.WriteString(controlCharacter)

	return fiscalCode.String(), nil
}

func reduce[T, M any](s []T, initValue M, op func(M, T) M) M {
	acc := initValue
	for _, v := range s {
		acc = op(acc, v)
	}
	return acc
}
