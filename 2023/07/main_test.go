package main

import (
	"reflect"
	"testing"
)

func Test_cardToSlice(t *testing.T) {
	type args struct {
		cards  string
		jokers bool
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{"32T3K", false},
			want: []int{3, 2, 10, 3, 13},
		},
		{
			args: args{"T55J5", false},
			want: []int{10, 5, 5, 11, 5},
		},
		{
			args: args{"T55J5", true},
			want: []int{10, 5, 5, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.cards, func(t *testing.T) {
			if got := cardToSlice(tt.args.cards, tt.args.jokers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cardToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_Rank(t *testing.T) {
	type fields struct {
		cardlist string
		jokers   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"5k",
			fields{
				"KKKKK",
				false,
			},
			25,
		},
		{
			"FH",
			fields{
				"KKQQQ",
				false,
			},
			13,
		},
		{
			"hc",
			fields{
				"12345",
				false,
			},
			5,
		},
		{
			"2 pair joker",
			fields{
				"22J66",
				true,
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				cardlist: tt.fields.cardlist,
				jokers:   tt.fields.jokers,
			}
			if got := h.Rank(); got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHands_WinningsOne(t *testing.T) {
	input := []string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}

	hands := makeHands(input, false)

	winnings := Winnings(hands)

	if winnings != 6440 {
		t.Errorf("want 6440, got %d", winnings)
	}
}

func TestHands_WinningsTwo(t *testing.T) {
	input := []string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}

	hands := makeHands(input, true)

	winnings := Winnings(hands)

	if winnings != 5905 {
		t.Errorf("want 5905, got %d", winnings)
	}
}
