package deck

import (
	"fmt"
	"testing"
	"time"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Joker})
	fmt.Println(Card{Rank: Three, Suit: Spade})

	//Output:
	// Ace of Hearts
	// Joker
	// Three of Spades
}
func TestNew(t *testing.T) {
	deck := New()
	if len(deck)!=52{
		t.Error("Not match the size")
	}
}
func TestSorting(t *testing.T){
	deck := New(NormalSort)
	card := Card{Rank: Ace, Suit: Spade}
	if deck[0]!=card{
		t.Error("Didn't sort")
	}
}


func TestShuffle(t *testing.T){
	deck := New(NormalSort)
	cards := New(NormalSort)
	deck.Shuffle()
	copy(cards,deck)
	time.Sleep(1e+9)
	cards.Shuffle()
	var counter int
	for i :=0;i<len(deck);i++{
		if deck[i]==cards[i]{
			counter++
		}
	}
	if counter == 52 {
	t.Error("Didn't shuffled")
	}
}


func TestJokers(t *testing.T){
	deck := New(NormalSort)
	deck.Jokers(4)
	if len(deck)!=56 {
		t.Error("Didn't add Jokers")
	}
}

func TestFilter(t *testing.T){
	filter := func (card Card) bool {
		return card.Suit==Spade || card.Rank == Seven
	}
	deck := New(NormalSort)
	deck.Filter(filter)
	if len(deck)==52 {
		t.Error("didn't filtered")
	}
}

func TestMoreDeck(t *testing.T){
	deck := New(NormalSort)
	deck = MoreDecks(deck, 4)
	if len(deck)!=(13*4*4) {
		t.Error("Didn't add correctly")
	}
}