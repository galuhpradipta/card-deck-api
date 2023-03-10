package main

import (
	"errors"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/galuhpradipta/card-deck-api/handlers"
	"github.com/galuhpradipta/card-deck-api/repositories"
	"github.com/galuhpradipta/card-deck-api/services"
)

var decks = make(map[string]Deck)

func main() {
	deckRepository := repositories.NewDeckRepository()
	deckService := services.NewDeckService(deckRepository)
	deckHandler := handlers.NewDeckHandler(deckService)

	app := fiber.New()

	app.Get("/decks/:id", getDeck)
	app.Post("/decks", deckHandler.CreateDeck)
	app.Post("/decks/:id/draw", drawDeckHandler)
	app.Listen(":3000")
}

var FullCardDecks = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS",
}

type Cards []string

type Card struct {
	Code  string `json:"code"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Deck struct {
	ID        string   `json:"id"`
	Shuffled  bool     `json:"shuffled"`
	Remaining int      `json:"remaining"`
	Cards     []string `json:"cards"`
}

type createDeckRequest struct {
	Shuffled bool     `json:"shuffled"`
	Cards    []string `json:"cards"`
}

func createDeckHandler(c *fiber.Ctx) error {
	var req createDeckRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	cards := Cards(FullCardDecks[:])
	if len(req.Cards) > 0 {
		cards = req.Cards
		err := cards.validate()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if req.Shuffled {
		cards = cards.shuffle()
	}

	id := uuid.New().String()
	deck := Deck{
		ID:        id,
		Shuffled:  false,
		Remaining: len(cards),
		Cards:     cards,
	}
	decks[deck.ID] = deck

	return c.JSON(deck)
}

type drawDeckRequest struct {
	Count int `json:"count"`
}

func drawDeckHandler(c *fiber.Ctx) error {
	var req drawDeckRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	if req.Count < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Count must be greater than 0",
		})
	}

	id := c.Params("id")
	deck, ok := decks[id]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Deck not found",
		})
	}

	if req.Count > deck.Remaining {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Count must be less than or equal to remaining cards",
		})
	}

	cards := deck.Cards[:req.Count]
	deck.Cards = deck.Cards[req.Count:]
	deck.Remaining = len(deck.Cards)
	decks[deck.ID] = deck

	var responseCards []Card
	for _, val := range cards {
		card := Card{
			Code:  val,
			Type:  cardTypeMap[val[1:]],
			Value: cardValueMap[val[:1]],
		}
		responseCards = append(responseCards, card)
	}

	return c.JSON(responseCards)
}

func (c Cards) validate() error {
	cards := Cards(FullCardDecks[:])
	cardPool := make(map[string]bool)
	for _, val := range c {
		if _, ok := cardPool[val]; ok {
			return errors.New("duplicate card: " + val)
		}

		if !cards.contain(val) {
			return errors.New("invalid card: " + val)
		}

		cardPool[val] = true
	}

	if len(c) > 52 {
		return errors.New("too many cards")
	}

	return nil
}

func (c Cards) contain(card string) bool {
	for _, val := range c {
		if val == card {
			return true
		}
	}
	return false
}

func (c Cards) shuffle() Cards {
	//TODO: add seed
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	return c
}

func getDeck(c *fiber.Ctx) error {
	id := c.Params("id")
	deck, ok := decks[id]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Deck not found",
		})
	}

	var cards []Card
	for _, val := range deck.Cards {
		card := Card{
			Code:  val,
			Type:  cardTypeMap[val[1:]],
			Value: cardValueMap[val[:1]],
		}
		cards = append(cards, card)
	}

	return c.JSON(GetDeckResponse{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
		Cards:     cards,
	})
}

type GetDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

var cardTypeMap = map[string]string{
	"C": "Clubs",
	"D": "Diamonds",
	"H": "Hearts",
	"S": "Spades",
}

var cardValueMap = map[string]string{
	"A": "Ace",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"T": "10",
	"J": "Jack",
	"Q": "Queen",
	"K": "King",
}
