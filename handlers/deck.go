package handlers

import (
	"errors"

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
	if err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
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

func (h *deckHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Deck ID is required",
		})
	}

	deck, err := h.deckService.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(deck)
}

func (h *deckHandler) Draw(c *fiber.Ctx) error {
	var req drawDeckRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if req.Count < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Count must be greater than 0",
		})
	}

	id := c.Params("id")
	cards, err := h.deckService.Draw(id, req.Count)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cards)
}

type drawDeckRequest struct {
	Count int `json:"count"`
}

type createDeckRequest struct {
	Shuffled bool     `json:"shuffled"`
	Cards    []string `json:"cards"`
}
