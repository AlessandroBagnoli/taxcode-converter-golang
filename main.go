package main

import (
	_ "embed"
	"github.com/sirupsen/logrus"
	_ "taxcode-converter/docs"
	"taxcode-converter/service/handler"
	"time"
)

var start = time.Now()

//go:embed assets/italian-cities.csv
var csvFile []byte

// @title taxcode-converter
// @version 1.0
// @host localhost:8080
// @BasePath /
//
//go:generate swag init --pd
func main() {
	h := handler.InitDependencies(csvFile)
	app := handler.CreateFiberApp(h)
	logrus.Infof("App started in %v", time.Since(start))
	logrus.Fatal(app.Listen(":8080"))
}
