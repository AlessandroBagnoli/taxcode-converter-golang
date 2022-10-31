package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mvmaasakkers/go-problemdetails"
	"taxcode-converter/service"
)

type Handler struct {
	service service.TaxCodeService
}

func NewHandler(service service.TaxCodeService) Handler {
	return Handler{service}
}

// CalculateTaxCode godoc
// @Summary Calculate tax code starting from the data of a person.
// @Accept json
// @Produce json
// @Param CalculateTaxCodeRequest body service.CalculateTaxCodeRequest true "CalculateTaxCodeRequest"
// @Success 200 {object} service.CalculateTaxCodeResponse
// @Failure 400 {object} problemdetails.ProblemDetails
// @Failure 404 {object} problemdetails.ProblemDetails
// @Failure 500 {object} problemdetails.ProblemDetails
// @Router /api/v1/taxcode:calculate-tax-code [post]
func (h Handler) CalculateTaxCode(c *fiber.Ctx) error {

	req := new(service.CalculateTaxCodeRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := h.service.CalculateTaxCode(*req)
	if err != nil {
		return err
	}

	return c.JSON(data)
}

// CalculatePersonData godoc
// @Summary Calculate data of a person starting from his tax code.
// @Accept json
// @Produce json
// @Param CalculatePersonDataRequest body service.CalculatePersonDataRequest true "CalculatePersonDataRequest"
// @Success 200 {object} service.CalculatePersonDataResponse
// @Failure 400 {object} problemdetails.ProblemDetails
// @Failure 404 {object} problemdetails.ProblemDetails
// @Failure 500 {object} problemdetails.ProblemDetails
// @Router /api/v1/taxcode:calculate-person-data [post]
func (h Handler) CalculatePersonData(c *fiber.Ctx) error {

	req := new(service.CalculatePersonDataRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := h.service.CalculatePersonData(*req)
	if err != nil {
		return err
	}

	return c.JSON(data)
}

// HandleError handles all the error wrapping them into a proper problemdetails.ProblemDetails object
func (h Handler) HandleError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()

	switch err.(type) {
	case *fiber.Error:
		var fiberError *fiber.Error
		errors.As(err, &fiberError)
		code = fiberError.Code
		message = fiberError.Message
	}

	problemDetails := problemdetails.New(code, "", "", message, c.Path())
	return c.Status(code).JSON(problemDetails)
}
