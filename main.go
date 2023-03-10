package main

import (
	"math/rand"

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

func createDecHandler(c *fiber.Ctx) error {
	var req createDeckRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	cards := Cards(FullCardDecks[:])
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
