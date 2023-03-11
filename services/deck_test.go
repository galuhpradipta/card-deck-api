package services

import (
	"testing"

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
