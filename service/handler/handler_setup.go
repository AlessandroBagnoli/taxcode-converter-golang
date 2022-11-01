package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swaggerfiber "github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"taxcode-converter/service/csv"
	"taxcode-converter/service/taxcode"
)

// InitDependencies creates and injects dependencies, returns the handler for incoming http requests
func InitDependencies(file []byte) Handler {
	v := taxcode.CreateValidator()
	p := csv.NewProcessor(file)
	t := taxcode.NewTaxCodeService(v, p)
	h := NewHandler(t)
	return h
}

// CreateFiberApp creates and configure the fiber server
func CreateFiberApp(h Handler) *fiber.App {
	// creation of fiber app with custom config for error handling, usage of logger and cors
	app := fiber.New(fiber.Config{ErrorHandler: h.HandleError})
	app.Use(configureFiberLogger())
	app.Use(cors.New())

	// routing for swagger documentation
	app.Get("/swagger/*", swaggerfiber.HandlerDefault)

	// routing for apis
	v1 := app.Group("/api/v1")
	v1.Post("/taxcode:calculate-tax-code", h.CalculateTaxCode)
	v1.Post("/taxcode:calculate-person-data", h.CalculatePersonData)

	// handling of signals for graceful shutdown
	gracefulShutDown(app)

	return app
}

func gracefulShutDown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			log.Infof("Gracefully shutting down...")
			_ = app.Shutdown()
		}
	}()
}

func configureFiberLogger() fiber.Handler {
	return logger.New(logger.Config{
		// this is because I just want to log incoming request to the exposed api and ignore everything else
		Next: func(c *fiber.Ctx) bool {
			return !strings.Contains(c.Path(), "/api/v1/")
		},
		Format: "[${time}]|${status}|${resBody}|${latency} - ${method}|${path}|${body}\n",
	})
}
