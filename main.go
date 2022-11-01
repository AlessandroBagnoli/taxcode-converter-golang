package main

import (
	_ "embed"
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
	_ = handler.CreateFiberApp(h).Listen(":8080")
}
