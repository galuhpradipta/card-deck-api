package repositories

import (
	"testing"

	"github.com/galuhpradipta/card-deck-api/shared"
	"github.com/stretchr/testify/assert"
)

var (
	partialDeck = []string{"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS"}
)

func Test_deckRepository_Create(t *testing.T) {
	type test struct {
		name     string
		pool     []string
		shuffled bool
		wantPool []string
	}
	tests := []test{
		{
			name:     "Create full deck",
			pool:     shared.FullCardDecks,
			shuffled: false,
			wantPool: shared.FullCardDecks,
		},
		{
			name:     "Create partial deck",
			pool:     partialDeck,
			shuffled: true,
			wantPool: partialDeck,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDeckRepository(shared.FullCardDecks)
			deck := r.Create(tt.pool, tt.shuffled)
			assert.NotEmpty(t, deck.DeckID)
			assert.Equal(t, tt.shuffled, deck.Shuffled)
			assert.Equal(t, len(tt.wantPool), len(deck.Pool))
		})
	}
}

func Test_deckRepository_GetDeck(t *testing.T) {
	type test struct {
		name      string
		deckID    string
		wantDeck  shared.Deck
		wantError string
	}
	tests := []test{
		{
			name:   "Get deck with valid deckID",
			deckID: mockTestID,
			wantDeck: shared.Deck{
				DeckID:    mockTestID,
				Shuffled:  false,
				Pool:      shared.FullCardDecks,
				Remaining: len(shared.FullCardDecks),
			},
			wantError: "",
		},
		{
			name:      "Get deck with invalid deckID",
			deckID:    "2",
			wantDeck:  shared.Deck{},
			wantError: ErrDeckNotFound.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMockDeckRepository()
			deck, err := r.GetDeck(tt.deckID)
			if err != nil {
				assert.Equal(t, tt.wantError, err.Error())
			} else {
				assert.Equal(t, tt.wantDeck, deck)
			}
		})
	}
}

func Test_deckRepository_Update(t *testing.T) {
	type test struct {
		name     string
		deckID   string
		pool     []string
		shuffled bool
		wantPool []string
		wantErr  string
	}
	tests := []test{
		{
			name:     "Update deck with valid deckID",
			deckID:   mockTestID,
			pool:     partialDeck,
			shuffled: true,
			wantPool: partialDeck,
			wantErr:  "",
		},
		{
			name:     "Update deck with invalid deckID",
			deckID:   "test-invalid-id",
			pool:     partialDeck,
			shuffled: true,
			wantPool: partialDeck,
			wantErr:  ErrDeckNotFound.Error(),
		},
		{
			name:     "Update deck with empty pool",
			deckID:   mockTestID,
			pool:     []string{},
			shuffled: true,
			wantPool: []string{},
			wantErr:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMockDeckRepository()
			deck, err := r.Update(tt.deckID, tt.pool, tt.shuffled)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			} else {
				assert.Equal(t, tt.deckID, deck.DeckID)
				assert.Equal(t, tt.shuffled, deck.Shuffled)
				assert.Equal(t, tt.wantPool, deck.Pool)
			}

		})
	}
}
