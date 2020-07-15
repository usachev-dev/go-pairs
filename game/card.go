package game

import (
	"image/color"
	"math/rand"
	"strconv"
	"strings"
)

const (
	Ace= iota
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

func allValues() []int {
	result := []int{}
	for i := Ace; i <= King; i++ {
		result = append(result, i)
	}
	return result
}

const (
	Hearts = iota
	Clubs
	Diamonds
	Spades

)

func allKinds() []int {
	result := []int{}
	for i := Hearts; i <= Spades; i++ {
		result = append(result, i)
	}
	return result
}


type Card struct {
	Value int
	Kind  int
}

func (c Card) DispayName() string {
	var value string
	switch c.Value {
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
	case Queen:
		value = "Queen"
	case King:
		value = "King"
	case Ace:
		value = "Ace"
	}

	var kind string
	switch c.Kind {
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

func (c Card) Unicode() string {
	offset := c.Kind* 13 + c.Value
	if offset > 25 {
		offset += 6
	}
	return string(65 + offset)
}

func (c Card) CardBack() string {
	return CardBack()
}

func NewCard(value int, kind int) Card {
	return Card {
		value, kind,
	}
}

func CardBack() string {
	return "0"
}

func RandCard() Card {
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
		Value: int(value),
		Kind:  int(kind),
	}
	return card, nil
}

func (c Card) Equals(card Card) bool {
	return c.Value == card.Value && c.Kind == card.Kind;
}

func (c Card) Color() color.Color {
	if c.Kind == Diamonds || c.Kind == Hearts {
		return color.RGBA{200,50,50, 255}
	} else {
		return color.RGBA{0,0,0,255}
	}
}

func (c Card ) BackColor() color.Color {
	return color.RGBA{0,0,0,255}
}