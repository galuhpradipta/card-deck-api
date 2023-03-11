package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/galuhpradipta/card-deck-api/repositories"
	"github.com/galuhpradipta/card-deck-api/shared"
)

type deckService struct {
	deckRepository repositories.DeckRepository
}

func NewDeckService(deckRepository repositories.DeckRepository) DeckService {
	return &deckService{
		deckRepository: deckRepository,
	}
}

func (s *deckService) Create(shuffled bool, cards []string) (shared.Deck, error) {
	deckCards := s.deckRepository.GetFullCardDecks()
	if len(cards) > 0 {
		deckCards = cards
		err := s.validate(deckCards)
		if err != nil {
			return shared.Deck{}, err
		}
	}

	if shuffled {
		deckCards = shuffle(deckCards)
	}

	deck := s.deckRepository.Create(deckCards, shuffled)
	return shared.Deck{
		DeckID:    deck.DeckID,
		Shuffled:  shuffled,
		Remaining: len(deck.Pool),
	}, nil
}

func (s *deckService) GetByID(id string) (shared.Deck, error) {
	deck, err := s.deckRepository.GetDeck(id)
	if err != nil {
		return shared.Deck{}, err
	}

	for _, v := range deck.Pool {
		deck.Cards = append(deck.Cards, shared.Card{
			Value: toValue(v[:len(v)-1]),
			Suit:  toSuit(v[len(v)-1:]),
			Code:  v,
		})
	}

	return deck, nil
}

func (s *deckService) Draw(id string, count int) ([]shared.Card, error) {
	deck, err := s.deckRepository.GetDeck(id)
	if err != nil {
		return []shared.Card{}, err
	}

	if count > len(deck.Pool) {
		return []shared.Card{}, errors.New("Count must be less than or equal to remaining cards")
	}

	drawedCards := deck.Pool[:count]
	deck.Pool = deck.Pool[count:]
	deck.Remaining = len(deck.Pool)
	s.deckRepository.Update(deck.DeckID, deck.Pool, deck.Shuffled)

	for _, v := range drawedCards {
		deck.Cards = append(deck.Cards, shared.Card{
			Value: toValue(v[:len(v)-1]),
			Suit:  toSuit(v[len(v)-1:]),
			Code:  v,
		})
	}

	return deck.Cards, nil
}

func (s *deckService) validate(cards []string) error {
	cardPool := make(map[string]bool)
	fullCardDecks := s.deckRepository.GetFullCardDecks()
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

func shuffle(cards []string) []string {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func contain(cards []string, card string) bool {
	for _, value := range cards {
		if value == card {
			return true
		}
	}
	return false
}

func toSuit(s string) string {
	return suitMap[s]
}

func toValue(s string) string {
	return valueMap[s]
}

var suitMap = map[string]string{
	"C": "CLUBS",
	"D": "DIAMONDS",
	"H": "HEARTS",
	"S": "SPADES",
}

var valueMap = map[string]string{
	"A": "Ace",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"T": "10",
	"J": "JACK",
	"Q": "QUEEN",
	"K": "KING",
}
