package cards

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

type Deck []Card

func NewDeck() Deck {
	deck := []Card{}
	for _, kind := range allKinds() {
		for _, value := range allValues() {
			deck = append(deck, Card{value, kind})
		}
	}
	return deck
}

func NewDeck36() Deck {
	deck := []Card{}
	for _, kind := range allKinds() {
		for _, value := range allValues()[4:] {
			deck = append(deck, Card{value, kind})
		}
	}
	return deck
}

func (d Deck) Shuffle() Deck {
	var newDeck Deck = []Card{}
	for range d {
		index := rand.Intn(len(d))
		newDeck = append(newDeck, d[index])
	}
	return newDeck
}

func (d Deck) Print() {
	for _, card := range d {
		fmt.Println(card.DispayName())
	}
}

func (d Deck) DrawCards(n int) (Deck, Deck) {
	// 1st are drawed cards
	// 2nd are remaining
	return d[:n], d[n:]
}

func (d Deck) ToString() string {
	result := ""
	for _, card := range d {
		result = result + fmt.Sprintf("%d %d,", card.value, card.kind)
	}
	return result
}

func nilDeck() Deck {
	return []Card{}
}

func deckFromString(str string) (Deck, error) {
	var strArr []string = strings.Split(str, ",")
	var deck Deck = []Card{}
	for _, s := range strArr {
		if len(s) == 0 {
			continue
		}
		card, err := cardFromString(s)
		if err != nil {
			return nilDeck(), err
		}
		deck = append(deck, card)
	}
	return deck, nil
}

func DeckFromFile(filename string) (Deck, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Card{}, err
	}
	return deckFromString(string(file))
}

func (d Deck) WriteToFile(filename string) error {
	str := d.ToString()
	return ioutil.WriteFile(filename, []byte(str), 0777)
}

func (d Deck) PrintDisplay() {
	for _, card := range d {
		fmt.Println(card.DispayName())
	}
}
