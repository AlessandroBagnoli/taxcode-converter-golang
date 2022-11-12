package taxcode

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
	"taxcode-converter/service"
	"time"
)

func calculate(req service.CalculateTaxCodeRequest, cityExtractor func(place service.Place) *service.CityCSV) (*service.CalculateTaxCodeResponse, error) {
	var fiscalCode bytes.Buffer
	fcSurname := strings.ToUpper(req.Surname) //TODO fine tune
	fcName := strings.ToUpper(req.Name)       // TODO fine tune
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
		return nil, service.NewCityNotPresentError(fmt.Sprintf("The city %s and province %s do not exixt", cityName, province))
	}
	fiscalCode.WriteString(city.Code)

	// control char
	evenSum := 0 //TODO complete
	oddSum := 0  // TODO complete
	controlInteger := (evenSum + oddSum) % 26
	controlCharacter := controlCharMap[controlInteger]
	fiscalCode.WriteString(controlCharacter)

	return &service.CalculateTaxCodeResponse{TaxCode: fiscalCode.String()}, nil
}

func consonants(word string) string {
	var consonants bytes.Buffer
	for _, char := range word {
		if !slices.Contains(vowelsSlice, char) {
			consonants.WriteRune(char)
		}
	}
	return consonants.String()
}

func vowels(word string) string {
	var consonants bytes.Buffer
	for _, char := range word {
		if slices.Contains(vowelsSlice, char) {
			consonants.WriteRune(char)
		}
	}
	return consonants.String()
}
