package repositories

import (
	"github.com/galuhpradipta/card-deck-api/shared"
	"github.com/google/uuid"
)

var fullCardDecks = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS",
}

type deckRepository struct {
	Decks map[string]shared.Deck
}

func NewDeckRepository() DeckRepository {
	return &deckRepository{
		Decks: make(map[string]shared.Deck),
	}
}

func (r *deckRepository) Create(pool []string, shuffled bool) string {
	deck := shared.Deck{
		ID:        r.generateID(),
		Shuffled:  shuffled,
		Remaining: len(pool),
		Pool:      pool,
	}

	r.Decks[deck.ID] = deck
	return deck.ID
}

// func (r *DeckRepository) GetDeck(id string) (shared.Deck, error) {
// 	deck, ok := r.Decks[id]
// 	if !ok {
// 		return deck, shared.ErrDeckNotFound
// 	}
// 	return deck, nil
// }

// func (r *DeckRepository) Draw(id string, count int) (shared.Deck, bool) {
// 	deck, ok := r.Decks[id]
// 	if !ok {
// 		return deck, false
// 	}

// 	if len(deck.Cards) < count {
// 		return deck, false
// 	}

// 	deck.Cards = deck.Cards[:count]
// 	deck.Remaining = len(deck.Cards)

// 	return deck, true
// }

func (r *deckRepository) generateID() string {
	return uuid.New().String()
}
