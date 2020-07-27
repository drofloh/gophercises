package main

import (
	"fmt"
	"strings"

	deck "github.com/drofloh/gophercises/deckofcards"
)

// Hand ...
type Hand []deck.Card

// String ...
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range strs {

		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString ...
func (h Hand) DealerString() string {
	return h[0].String() + ", *HIDDEN*"
}

// Score ...
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// MinScore ...
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Shuffle ...
func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

// Deal ...
func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

// Stand ...
func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

// Hit ...
func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

// EndHand ...
func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("FINAL HANDS")

	fmt.Println("Player:", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("Bust")
	case dScore > 21:
		fmt.Println("Dealer Bust")
	case pScore > dScore:
		fmt.Println("You Win")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()

	ret.Player = nil
	ret.Dealer = nil

	return ret
}
func main() {
	var gs GameState
	gs = Shuffle(gs)

	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		// cards := deck.New(deck.Deck(3), deck.Shuffle)

		// var gs GameState
		// gs.Deck = deck.New(deck.Deck(3), deck.Shuffle)

		// var card deck.Card
		// //for i := 0; i < 10; i++ {
		// //	card, cards = cards[0], cards[1:]
		// //	fmt.Println(card)
		// //}
		// //var h Hand = cards[0:3]
		// //fmt.Println(h)
		// var player, dealer Hand
		// for i := 0; i < 2; i++ {
		// 	for _, hand := range []*Hand{&player, &dealer} {
		// 		card, cards = draw(cards)
		// 		*hand = append(*hand, card)
		// 	}
		// }
		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player: ", gs.Player)
			fmt.Println("Dealer: ", gs.Dealer.DealerString())
			fmt.Println("What next? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		// for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		// 	card, cards = draw(cards)
		// 	dealer = append(dealer, card)
		// }

		gs = EndHand(gs)
	}
	// pScore, dScore := player.Score(), dealer.Score()
	// fmt.Println("FINAL HANDS")

	// fmt.Println("Player:", player, "\nScore:", pScore)
	// fmt.Println("Dealer:", dealer, "\nScore:", dScore)
	// switch {
	// case pScore > 21:
	// 	fmt.Println("Bust")
	// case dScore > 21:
	// 	fmt.Println("Dealer Bust")
	// case pScore > dScore:
	// 	fmt.Println("You Win")
	// case dScore > pScore:
	// 	fmt.Println("You lose")
	// case dScore == pScore:
	// 	fmt.Println("Draw")
	// }

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// State ...
type State int8

// Set some states ...
const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// GameState ...
type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

// CurrentPlayer ...
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It isnt currently any players turn")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
