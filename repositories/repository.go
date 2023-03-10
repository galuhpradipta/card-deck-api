package repositories

type DeckRepository interface {
	Create(pool []string, shuffled bool) string
}
