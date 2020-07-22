package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Four, Suit: Diamond})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Four of Diamonds
	// Nine of Clubs
	// Queen of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	card := Card{Suit: Spade, Rank: Ace}
	if cards[0] != card {
		t.Error("expected the ace of spades as first card, got:", cards[0])
	}

}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	card := Card{Suit: Spade, Rank: Ace}
	if cards[0] != card {
		t.Error("expected the ace of spades as first card, got:", cards[0])
	}

}

func TestJoker(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {

		t.Error("Expeted 3 jokers, got: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("All 2 and 3 cards should be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 52*3 {
		t.Errorf("expectd %d cards, received %d cards", 52*3, len(cards))
	}
}
