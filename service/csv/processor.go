package csv

import (
	"embed"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"io"
	"taxcode-converter/service"
)

//go:embed assets/italian-cities.csv
var content embed.FS

type Processor struct{}

func NewProcessor() Processor {
	return Processor{}
}

func (Processor) ParseCities() []service.CityCSV {
	clientsFile, err := content.Open("assets/italian-cities.csv")
	if err != nil {
		log.Panic(err)
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.TrimLeadingSpace = true
		r.Comma = ';'
		return r
	})

	var cities []service.CityCSV
	if err = gocsv.Unmarshal(clientsFile, &cities); err != nil {
		log.Panic(err)
	}

	log.Infof("Loaded %v cities from csv file", len(cities))
	return cities
}

func (p Processor) GetCityCodesCache(cities []service.CityCSV) map[string]service.CityCSV {
	//TODO implement me
	panic("implement me")
}

func (p Processor) GetCityPlacesCache(cities []service.CityCSV) map[service.Place]service.CityCSV {
	//TODO implement me
	panic("implement me")
}
