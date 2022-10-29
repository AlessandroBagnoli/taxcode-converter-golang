package handler

import (
	"github.com/gofiber/fiber/v2"
	"taxcode-converter/service"
)

type Handler struct {
	service service.TaxCodeService
}

func NewHandler(service service.TaxCodeService) *Handler {
	return &Handler{service}
}

// CalculateTaxCode godoc
// @Summary Calculate tax code starting from the data of a person.
// @Accept json
// @Produce json
// @Param CalculateTaxCodeRequest body service.CalculateTaxCodeRequest true "CalculateTaxCodeRequest"
// @Success 200 {object} service.CalculateTaxCodeResponse
// @Router /api/v1/taxcode:calculate-tax-code [post]
func (h *Handler) CalculateTaxCode(c *fiber.Ctx) error {

	req := new(service.CalculateTaxCodeRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, _ := h.service.CalculateTaxCode(c.Context(), req)
	return c.JSON(data)
}

// CalculatePersonData godoc
// @Summary Calculate data of a person starting from his tax code.
// @Accept json
// @Produce json
// @Param CalculatePersonDataRequest body service.CalculatePersonDataRequest true "CalculatePersonDataRequest"
// @Success 200 {object} service.CalculatePersonDataResponse
// @Router /api/v1/taxcode:calculate-person-data [post]
func (h *Handler) CalculatePersonData(c *fiber.Ctx) error {

	req := new(service.CalculatePersonDataRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, _ := h.service.CalculatePersonData(c.Context(), req)
	return c.JSON(data)
}
