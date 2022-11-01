package main

import (
	_ "embed"
	log "github.com/sirupsen/logrus"
	_ "taxcode-converter/docs"
	"taxcode-converter/service/handler"
)

// @title taxcode-converter
// @version 1.0
// @host localhost:8080
// @BasePath /
//
//go:generate swag init --pd

//go:embed assets/italian-cities.csv
var csvFile []byte

func main() {
	h := handler.InitDependencies(csvFile)
	app := handler.CreateFiberApp(h)
	log.Fatal(app.Listen(":8080"))
}
