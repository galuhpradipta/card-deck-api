package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var decks = make(map[string]Deck)

func main() {
	app := fiber.New()
	app.Post("/decks", createDecHandler)
	app.Get("/decks/:id", getDeck)
	app.Listen(":3000")
}

var FullCardDecks = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
}

type Cards []string

type Deck struct {
	ID        string `json:"id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func createDecHandler(c *fiber.Ctx) error {
	cards := FullCardDecks
	id := uuid.New().String()
	deck := Deck{
		ID:        id,
		Shuffled:  false,
		Remaining: len(cards),
	}
	decks[deck.ID] = deck

	return c.JSON(deck)
}

func getDeck(c *fiber.Ctx) error {
	id := c.Params("id")
	deck, ok := decks[id]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Deck not found",
		})
	}
	return c.JSON(deck)
}
