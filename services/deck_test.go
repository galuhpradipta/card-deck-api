package services

import (
	"testing"

	"github.com/galuhpradipta/card-deck-api/repositories"
	"github.com/galuhpradipta/card-deck-api/shared"
	"github.com/stretchr/testify/assert"
)

func Test_toSuit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test toSuit with valid suit HEARTS",
			args: args{
				s: "H",
			},
			want: "HEARTS",
		},
		{
			name: "Test toSuit with valid suit DIAMONDS",
			args: args{
				s: "D",
			},
			want: "DIAMONDS",
		},
		{
			name: "Test toSuit with valid suit SPADES",
			args: args{
				s: "S",
			},
			want: "SPADES",
		},
		{
			name: "Test toSuit with valid suit CLUBS",
			args: args{
				s: "C",
			},
			want: "CLUBS",
		},
		{
			name: "Test toSuit with invalid suit",
			args: args{
				s: "X",
			},
			want: "",
		},
		{
			name: "Test toSuit with empty suit",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSuit(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toValue(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test toValue with valid value ACE",
			args: args{
				s: "A",
			},
			want: "ACE",
		},
		{
			name: "Test toValue with valid value TWO",
			args: args{
				s: "2",
			},
			want: "2",
		},
		{
			name: "Test toValue with valid value THREE",
			args: args{
				s: "3",
			},
			want: "3",
		},
		{
			name: "Test toValue with valid value FOUR",
			args: args{
				s: "4",
			},
			want: "4",
		},
		{
			name: "Test toValue with valid value FIVE",
			args: args{
				s: "5",
			},
			want: "5",
		},
		{
			name: "Test toValue with valid value SIX",
			args: args{
				s: "6",
			},
			want: "6",
		},
		{
			name: "Test toValue with valid value SEVEN",
			args: args{
				s: "7",
			},
			want: "7",
		},
		{
			name: "Test toValue with valid value EIGHT",
			args: args{
				s: "8",
			},
			want: "8",
		},
		{
			name: "Test toValue with valid value NINE",
			args: args{
				s: "9",
			},
			want: "9",
		},
		{
			name: "Test toValue with valid value TEN",
			args: args{
				s: "T",
			},
			want: "10",
		},
		{
			name: "Test toValue with valid value JACK",
			args: args{
				s: "J",
			},
			want: "JACK",
		},
		{
			name: "Test toValue with valid value QUEEN",
			args: args{
				s: "Q",
			},
			want: "QUEEN",
		},
		{
			name: "Test toValue with valid value KING",
			args: args{
				s: "K",
			},
			want: "KING",
		},
		{
			name: "Test toValue with invalid value",
			args: args{
				s: "X",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toValue(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_contain(t *testing.T) {
	type args struct {
		cards []string
		card  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test contain with valid card",
			args: args{
				cards: []string{"2H", "3D", "4S", "5C"},
				card:  "4S",
			},
			want: true,
		},
		{
			name: "Test contain with invalid card",
			args: args{
				cards: []string{"2H", "3D", "4S", "5C"},
				card:  "4X",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contain(tt.args.cards, tt.args.card)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_shuffle(t *testing.T) {
	type args struct {
		cards []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test shuffle with valid cards",
			args: args{
				cards: []string{"2H", "3D", "4S", "5C"},
			},
			want: []string{"2H", "3D", "4S", "5C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shuffle(tt.args.cards)
			assert.Equal(t, len(tt.want), len(got))
			for i := range got {
				assert.Contains(t, tt.want, got[i])
			}
		})
	}
}

func Test_deckService_validate(t *testing.T) {
	type fields struct {
		deckRepository repositories.DeckRepository
	}
	type args struct {
		cards []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test validate with valid cards",
			fields: fields{
				deckRepository: repositories.NewMockDeckRepository(),
			},
			args: args{
				cards: []string{"2H", "3D", "4S", "5C"},
			},
			wantErr: false,
		},
		{
			name: "Test validate with invalid cards",
			fields: fields{
				deckRepository: repositories.NewMockDeckRepository(),
			},
			args: args{
				cards: []string{"2H", "3D", "4S", "5C", "6X"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &deckService{
				deckRepository: tt.fields.deckRepository,
			}
			err := s.validate(tt.args.cards)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_deckService_Create(t *testing.T) {
	type fields struct {
		deckRepository repositories.DeckRepository
	}
	type args struct {
		shuffled bool
		cards    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    shared.Deck
		wantErr bool
	}{
		{
			name: "Test Create with valid cards",
			fields: fields{
				deckRepository: repositories.NewDeckRepository(shared.FullCardDecks),
			},
			args: args{
				shuffled: false,
				cards:    []string{"2H", "3D", "4S", "5C"},
			},
			want: shared.Deck{
				DeckID:    repositories.MockTestID,
				Shuffled:  false,
				Remaining: 4,
				Cards: []shared.Card{
					{
						Value: "2",
						Suit:  "HEARTS",
						Code:  "2H",
					},
					{
						Value: "3",
						Suit:  "DIAMONDS",
						Code:  "3D",
					},
					{
						Value: "4",
						Suit:  "SPADES",
						Code:  "4S",
					},
					{
						Value: "5",
						Suit:  "CLUBS",
						Code:  "5C",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &deckService{
				deckRepository: tt.fields.deckRepository,
			}
			got, err := s.Create(tt.args.shuffled, tt.args.cards)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got.DeckID)
				assert.Equal(t, tt.want.Shuffled, got.Shuffled)
				assert.Equal(t, tt.want.Remaining, got.Remaining)
				assert.Equal(t, len(tt.want.Pool), len(got.Pool))
				for i := range got.Cards {
					assert.Equal(t, tt.want.Cards[i].Value, got.Cards[i].Value)
					assert.Equal(t, tt.want.Cards[i].Suit, got.Cards[i].Suit)
					assert.Equal(t, tt.want.Cards[i].Code, got.Cards[i].Code)
				}
			}
		})
	}
}

func Test_deckService_GetByID(t *testing.T) {
	type fields struct {
		deckRepository repositories.DeckRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    shared.Deck
		wantErr string
	}{
		{
			name: "Test GetByID with valid id",
			fields: fields{
				deckRepository: repositories.NewMockDeckRepository(),
			},
			args: args{
				id: repositories.MockTestID,
			},
			want: shared.Deck{
				DeckID:    repositories.MockTestID,
				Shuffled:  false,
				Remaining: 52,
				Pool:      shared.FullCardDecks,
			},
			wantErr: "",
		},
		{
			name: "Test GetByID with invalid id",
			fields: fields{
				deckRepository: repositories.NewMockDeckRepository(),
			},
			args: args{
				id: "invalid-id",
			},
			want:    shared.Deck{},
			wantErr: repositories.ErrDeckNotFound.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &deckService{
				deckRepository: tt.fields.deckRepository,
			}
			got, err := s.GetByID(tt.args.id)
			if tt.wantErr != "" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.DeckID, got.DeckID)
				assert.Equal(t, tt.want.Shuffled, got.Shuffled)
				assert.Equal(t, tt.want.Remaining, got.Remaining)
				assert.Equal(t, len(tt.want.Pool), len(got.Pool))
			}
		})
	}
}

func Test_deckService_Draw(t *testing.T) {
	type fields struct {
		deckRepository repositories.DeckRepository
	}
	type args struct {
		id    string
		count int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantDrawn int
		wantErr   string
	}{
		{
			name: "Test Draw with valid id and count",
			fields: fields{
				deckRepository: repositories.NewMockDeckRepository(),
			},
			args: args{
				id:    repositories.MockTestID,
				count: 2,
			},
			wantDrawn: 2,
			wantErr:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &deckService{
				deckRepository: tt.fields.deckRepository,
			}
			got, err := s.Draw(tt.args.id, tt.args.count)
			if tt.wantErr != "" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantDrawn, len(got))
			}
		})
	}
}
