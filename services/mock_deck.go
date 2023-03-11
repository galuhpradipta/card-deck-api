package services

import "github.com/galuhpradipta/card-deck-api/shared"

type mockDeckService struct {
}

func NewMockDeckService() DeckService {
	return &mockDeckService{}
}

func (m *mockDeckService) Create(shuffled bool, cards []string) (shared.Deck, error) {
	return shared.Deck{}, nil
}

func (m *mockDeckService) GetByID(id string) (shared.Deck, error) {
	return shared.Deck{}, nil
}

func (m *mockDeckService) Draw(id string, count int) ([]shared.Card, error) {
	return []shared.Card{}, nil
}
