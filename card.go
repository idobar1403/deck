//go:generate stringer -type=Rank,Suit
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank int

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

type Deck []Card

func New(sorts ...func(Deck) Deck) Deck {
	var deck Deck
	for i := Ace; i <= King; i++ {
		for j := Spade; j <= Heart; j++ {
			deck = append(deck, Card{Rank: i, Suit: j})
		}
	}
	for _, sortType := range sorts {
		deck = sortType(deck)
	}
	return deck
}
func NormalSort(cards Deck) Deck{
	sort.Slice(cards,Less(cards))
	return cards
}
func Less(cards Deck) func(i int, j int) bool {
	return func(i int, j int) bool {
		return getAbsRank(cards[i]) < getAbsRank(cards[j])
	}
}
func getAbsRank(card Card) int {
	return int(card.Suit)*int(King) + int(card.Rank)
}
func (cards *Deck) Shuffle(){
	shuffleDeck := make([]Card, len(*cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for ind, val := range r.Perm(52){
		shuffleDeck[ind] = (*cards)[val]
	}
	*cards = shuffleDeck
}
func (cards *Deck) Jokers(n int){
	for i:=0;i<n;i++ {
		*cards = append(*cards,Card{
			Rank: Rank(i),
			 Suit: Joker,
			})
	}
}

func (cards *Deck) Filter(filter func(card Card) bool){
	
		var filteredDeck Deck
		for _,val:=range *cards{
			if!filter(val){
				filteredDeck = append(filteredDeck, val)
			}
		}
		*cards = filteredDeck
}


func MoreDecks(cards Deck,n int) Deck{
	var newDeck Deck
	for i:=0;i<n;i++{
		newDeck = append(newDeck, cards...)
	}
	return newDeck
}