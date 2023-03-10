package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/galuhpradipta/card-deck-api/repositories"
	"github.com/galuhpradipta/card-deck-api/shared"
)

type DeckService struct {
	deckRepository *repositories.DeckRepository
}

func NewDeckService(deckRepository *repositories.DeckRepository) *DeckService {
	return &DeckService{
		deckRepository: deckRepository,
	}
}

func (s *DeckService) CreateDeck(shuffled bool, cards []string) (shared.Deck, error) {
	deckCards := fullCardDecks
	if len(cards) > 0 {
		deckCards = cards
		err := validate(deckCards)
		if err != nil {
			return shared.Deck{}, err
		}
	}

	if shuffled {
		deckCards = shuffle(deckCards)
	}

	deckID := s.deckRepository.CreateDeck(deckCards, shuffled)

	return shared.Deck{
		ID:        deckID,
		Shuffled:  shuffled,
		Remaining: len(deckCards),
	}, nil
}

var fullCardDecks = []string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS",
}

func shuffle(cards []string) []string {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func validate(cards []string) error {
	cardPool := make(map[string]bool)
	for _, val := range cards {
		if _, ok := cardPool[val]; ok {
			return errors.New("duplicate card: " + val)
		}

		if !contain(fullCardDecks, val) {
			return errors.New("invalid card: " + val)
		}

		cardPool[val] = true
	}

	return nil
}

func contain(cards []string, card string) bool {
	for _, value := range cards {
		if value == card {
			return true
		}
	}
	return false
}