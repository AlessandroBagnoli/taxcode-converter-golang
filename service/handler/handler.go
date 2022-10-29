package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-poc/service"
)

type Handler struct {
	service service.TaxCodeService
}

func NewHandler(service service.TaxCodeService) *Handler {
	return &Handler{service}
}

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
