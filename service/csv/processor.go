package csv

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"io"
	"taxcode-converter/service"
)

type Processor struct {
	cities          []*service.CityCSV
	cityCodesCache  map[string]*service.CityCSV
	cityPlacesCache map[service.Place]*service.CityCSV
}

func NewProcessor(file io.Reader) Processor {
	cities := parseCities(file)
	cityCodesCache := createCityCodesCache(cities)
	cityPlacesCache := createCityPlacesCache(cities)
	return Processor{
		cities:          cities,
		cityCodesCache:  cityCodesCache,
		cityPlacesCache: cityPlacesCache,
	}
}

func (p Processor) CityFromCode(code string) *service.CityCSV {
	return p.cityCodesCache[code]
}

func (p Processor) CityFromPlace(place service.Place) *service.CityCSV {
	return p.cityPlacesCache[place]
}

func parseCities(file io.Reader) []*service.CityCSV {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.TrimLeadingSpace = true
		r.Comma = ';'
		return r
	})

	var cities []*service.CityCSV
	if err := gocsv.Unmarshal(file, &cities); err != nil {
		log.Panic(err)
	}

	log.Infof("Loaded %v cities from csv file", len(cities))
	return cities
}

func createCityCodesCache(cities []*service.CityCSV) map[string]*service.CityCSV {
	cache := make(map[string]*service.CityCSV)
	for _, city := range cities {
		cache[city.Code] = city
	}
	return cache
}

func createCityPlacesCache(cities []*service.CityCSV) map[service.Place]*service.CityCSV {
	cache := make(map[service.Place]*service.CityCSV)
	for _, city := range cities {
		place := service.Place{
			CityName: city.Name,
			Province: city.Province,
		}
		cache[place] = city
	}
	return cache
}
