package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/sirupsen/logrus"
	_ "taxcode-converter/docs"
	"taxcode-converter/service/config"
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
	conf := config.NewConfig()
	h := handler.InitDependencies(bytes.NewReader(csvFile))
	app := handler.CreateFiberApp(h)
	logrus.Infof("App started in %f seconds", time.Since(start).Seconds())
	_ = app.Listen(fmt.Sprintf(":%d", conf.RestPort))
}
