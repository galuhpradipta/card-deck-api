package repositories

import (
	"github.com/galuhpradipta/card-deck-api/shared"
)

const (
	MockTestID = "test-id"
)

type mockDeckRepository struct {
	Decks         map[string]shared.Deck
	FullCardDecks []string
}

func NewMockDeckRepository() DeckRepository {
	mockDecks := make(map[string]shared.Deck)
	mockDecks[MockTestID] = shared.Deck{
		DeckID:    MockTestID,
		Shuffled:  false,
		Pool:      shared.FullCardDecks,
		Remaining: len(shared.FullCardDecks),
	}

	return &mockDeckRepository{
		Decks:         mockDecks,
		FullCardDecks: shared.FullCardDecks,
	}
}

func (r *mockDeckRepository) Create(pool []string, shuffled bool) shared.Deck {
	return r.Decks[MockTestID]
}

func (r *mockDeckRepository) GetDeck(id string) (shared.Deck, error) {
	deck, ok := r.Decks[id]
	if !ok {
		return deck, ErrDeckNotFound
	}
	deck.Remaining = len(deck.Pool)
	return deck, nil
}

func (r *mockDeckRepository) Update(id string, pool []string, shuffled bool) (shared.Deck, error) {
	if r.Decks[id].DeckID == "" {
		return shared.Deck{}, ErrDeckNotFound
	}

	deck := r.Decks[id]
	deck.Pool = pool
	deck.Shuffled = shuffled
	r.Decks[id] = deck
	return deck, nil
}

func (r *mockDeckRepository) GetFullCardDecks() []string {
	return r.FullCardDecks
}
