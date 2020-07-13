package cards

import (
	"math/rand"
	"strconv"
	"strings"
)

const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Dame
	King
	Ace
)

func allValues() []int {
	result := []int{}
	for i := Two; i <= Ace; i++ {
		result = append(result, i)
	}
	return result
}

const (
	Spades = iota
	Hearts
	Diamonds
	Clubs
)

func allKinds() []int {
	result := []int{}
	for i := Spades; i <= Clubs; i++ {
		result = append(result, i)
	}
	return result
}

type Card struct {
	value int
	kind  int
}

func (c Card) DispayName() string {
	var value string
	switch c.value {
	case Two:
		value = "Two"
	case Three:
		value = "Three"
	case Four:
		value = "Four"
	case Five:
		value = "Five"
	case Six:
		value = "Six"
	case Seven:
		value = "Seven"
	case Eight:
		value = "Eight"
	case Nine:
		value = "Nine"
	case Ten:
		value = "Ten"
	case Jack:
		value = "Jack"
	case Dame:
		value = "Dame"
	case King:
		value = "King"
	case Ace:
		value = "Ace"
	}

	var kind string
	switch c.kind {
	case Hearts:
		kind = "Hearts"
	case Diamonds:
		kind = "Diamonds"
	case Clubs:
		kind = "Clubs"
	case Spades:
		kind = "Spades"
	}

	var result string = value + " of " + kind
	return result
}

func randCard() Card {
	var value int = rand.Intn(Ace)
	var kind int = rand.Intn(4)
	var card Card = Card{
		value,
		kind,
	}
	return card
}


func cardFromString(s string) (Card, error) {
	split := strings.Split(s, " ")
	var v string = split[0]
	var k string = split[1]
	value, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return Card{}, err
	}
	kind, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return Card{}, err
	}
	var card Card = Card{
		value: int(value),
		kind:  int(kind),
	}
	return card, nil
}

func (c Card) equals(card Card) bool {
	return c.value == card.value && c.kind == card.kind;
}