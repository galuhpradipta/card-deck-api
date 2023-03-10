package services

import "github.com/galuhpradipta/card-deck-api/shared"

type DeckService interface {
	Create(shuffled bool, cards []string) (shared.Deck, error)
	GetByID(id string) (shared.Deck, error)
	Draw(id string, count int) ([]shared.Card, error)
}
