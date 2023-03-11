package repositories

import (
	"errors"

	"github.com/galuhpradipta/card-deck-api/shared"
	"github.com/google/uuid"
)

type deckRepository struct {
	Decks         map[string]shared.Deck
	FullCardDecks []string
}

func NewDeckRepository(fullCardDecks []string) DeckRepository {
	return &deckRepository{
		Decks:         make(map[string]shared.Deck),
		FullCardDecks: fullCardDecks,
	}
}

func (r *deckRepository) Create(pool []string, shuffled bool) string {
	deck := shared.Deck{
		DeckID:   r.generateID(),
		Shuffled: shuffled,
		Pool:     pool,
	}

	r.Decks[deck.DeckID] = deck
	return deck.DeckID
}

func (r *deckRepository) GetDeck(id string) (shared.Deck, error) {
	deck, ok := r.Decks[id]
	if !ok {
		return deck, errors.New("Deck not found")
	}
	deck.Remaining = len(deck.Pool)
	return deck, nil
}

func (r *deckRepository) Update(id string, pool []string, shuffled bool) {
	deck := r.Decks[id]
	deck.Pool = pool
	deck.Shuffled = shuffled
	r.Decks[id] = deck
}

func (r *deckRepository) GetFullCardDecks() []string {
	return r.FullCardDecks
}

func (r *deckRepository) generateID() string {
	return uuid.New().String()
}
