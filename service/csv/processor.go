package csv

import (
	"embed"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"io"
	"taxcode-converter/service"
)

// TODO I dont want the actual resource inside the module...find a way to externalize it and at the same time to keep the methods testable
//
//go:embed assets/italian-cities.csv
var content embed.FS

type Processor struct {
	cities          []service.CityCSV
	cityCodesCache  map[string]service.CityCSV
	cityPlacesCache map[service.Place]service.CityCSV
}

func (p Processor) CityCodesCache() map[string]service.CityCSV {
	return p.cityCodesCache
}

func (p Processor) CityPlacesCache() map[service.Place]service.CityCSV {
	return p.cityPlacesCache
}

func NewProcessor() Processor {
	cities := parseCities()
	cityCodesCache := createCityCodesCache(cities)
	cityPlacesCache := createCityPlacesCache(cities)
	return Processor{
		cities:          cities,
		cityCodesCache:  cityCodesCache,
		cityPlacesCache: cityPlacesCache,
	}
}

func parseCities() []service.CityCSV {
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

func createCityCodesCache(cities []service.CityCSV) map[string]service.CityCSV {
	cache := make(map[string]service.CityCSV)
	for _, city := range cities {
		cache[city.Code] = city
	}
	return cache
}

func createCityPlacesCache(cities []service.CityCSV) map[service.Place]service.CityCSV {
	cache := make(map[service.Place]service.CityCSV)
	for _, city := range cities {
		place := service.Place{
			CityName: city.Name,
			Province: city.Province,
		}
		cache[place] = city
	}
	return cache
}
