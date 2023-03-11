package shared

import "errors"

var FullCardDecks = []string{
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD",
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH",
}

var (
	ErrDeckNotFound = errors.New("deck not found")
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Deck struct {
	DeckID    string   `json:"deck_id"`
	Shuffled  bool     `json:"shuffled"`
	Remaining int      `json:"remaining"`
	Cards     []Card   `json:"cards,omitempty"`
	Pool      []string `json:"-"`
}
