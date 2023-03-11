package repositories

import "github.com/galuhpradipta/card-deck-api/shared"

type DeckRepository interface {
	GetDeck(id string) (shared.Deck, error)
	GetFullCardDecks() []string
	Create(pool []string, shuffled bool) shared.Deck
	Update(id string, pool []string, shuffled bool) (shared.Deck, error)
}
