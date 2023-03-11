package shared

import "errors"

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
