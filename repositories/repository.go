package repositories

import "github.com/galuhpradipta/card-deck-api/shared"

type DeckRepository interface {
	Create(pool []string, shuffled bool) string
	GetDeck(id string) (shared.Deck, error)
	Update(id string, pool []string, shuffled bool)
}
