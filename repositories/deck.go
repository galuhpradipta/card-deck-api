package repositories

import (
	"errors"

	"github.com/galuhpradipta/card-deck-api/shared"
	"github.com/google/uuid"
)

var (
	ErrDeckNotFound = errors.New("Deck not found")
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

func (r *deckRepository) Create(pool []string, shuffled bool) shared.Deck {
	deck := shared.Deck{
		DeckID:    r.generateID(),
		Shuffled:  shuffled,
		Remaining: len(pool),
		Pool:      pool,
	}

	r.Decks[deck.DeckID] = deck
	return deck
}

func (r *deckRepository) GetDeck(id string) (shared.Deck, error) {
	deck, ok := r.Decks[id]
	if !ok {
		return deck, ErrDeckNotFound
	}
	deck.Remaining = len(deck.Pool)
	return deck, nil
}

func (r *deckRepository) Update(id string, pool []string, shuffled bool) (shared.Deck, error) {
	if r.Decks[id].DeckID != id {
		return shared.Deck{}, ErrDeckNotFound
	}

	deck := r.Decks[id]
	deck.Pool = pool
	deck.Shuffled = shuffled
	r.Decks[id] = deck
	return deck, nil
}

func (r *deckRepository) GetFullCardDecks() []string {
	return r.FullCardDecks
}

func (r *deckRepository) generateID() string {
	return uuid.New().String()
}
