package handlers

import (
	"github.com/galuhpradipta/card-deck-api/services"
	"github.com/gofiber/fiber/v2"
)

type deckHandler struct {
	deckService services.DeckService
}

func NewDeckHandler(deckService services.DeckService) DeckHandler {
	return &deckHandler{
		deckService: deckService,
	}
}

func (h *deckHandler) Create(c *fiber.Ctx) error {
	var req createDeckRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	deck, err := h.deckService.Create(req.Shuffled, req.Cards)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(deck)
}

type createDeckRequest struct {
	Shuffled bool     `json:"shuffled"`
	Cards    []string `json:"cards"`
}
