package cards

import (
	"os"
	"testing"
)

func TestNewDeckSize(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 52 {
		t.Fail()
	}
}
func TestNewDeck36Size(t *testing.T) {
	deck := NewDeck36()
	if len(deck) != 36 {
		t.Fail()
	}
}
func TestShuffledDeckIsDifferentButSameLength(t *testing.T) {
	deck := NewDeck()
	shuffledDeck := deck.Shuffle()
	isSame := true
	if (len(deck) != len(shuffledDeck)) {
		t.Errorf("Deck and shuffled deck should have same size, got %v and %v", len(deck), len(shuffledDeck))
	}
	for i := range shuffledDeck {
		if deck[i].equals(shuffledDeck[i]) {
			isSame = false;
			break
		}
	}
	if isSame {
		t.Errorf("Deck and shuffled deck should be different")
	}
}
func TestDrawCardsHasCorrectLen(t *testing.T) {
	deck := NewDeck()
	hand, remaining := deck.DrawCards(5)
	if (len(hand) != 5) {
		t.Errorf("Expected draw 5 to draw 5, got %v", len(hand))
	}
	if (len(hand) + len(remaining) != len(deck)) {
		t.Errorf("Expected draw + remaining cards cound to be the same as deck %v, got %v", len(deck), len(hand) + len(remaining))
	}
}
func TestCanSaveAndLoadDeckToFile(t *testing.T) {
	deck := NewDeck()
	filename := "deck_testdata.txt"
	err := deck.WriteToFile(filename)
	defer os.Remove(filename)
	if (err != nil) {
		t.Errorf("Could not write deck to file")
	}
	_, er := DeckFromFile(filename)
	if (er != nil) {
		t.Errorf("Could not read written file to deck")
	}
}