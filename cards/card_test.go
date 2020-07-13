package cards

import "testing"

func TestValuesCount(t *testing.T) {
	values := allValues()
	if len(values) != 13 {
		t.Errorf("Expected card values count 13, got %v", len(values))
	}
}
func TestKindsCount(t *testing.T) {
	kinds := allKinds()
	if len(kinds) != 4 {
		t.Errorf("Expected card kinds count 4, got %v", len(kinds))
	}
}
func TestAllCardsHaveDisplayName(t *testing.T) {
	deck := NewDeck()
	for _, card := range deck {
		dispayName := card.DispayName()
		if len(dispayName) == 0 {
			t.Errorf("All cards should be able to display their names")
		}
		if len(dispayName) <= 3 {
			t.Errorf("All cards should display names should have 3 or move characters, was %v", dispayName)
		}
	}
}
func TestRandCardReturnsDifferentCards(t *testing.T) {
	cards := []Card{}
	for i :=0; i<=999; i++ {
		card := randCard();
		hasSame := false
		for _, c := range cards {
			if c.equals(card) {
				hasSame = true
			}
		}
		if  !hasSame {
			cards = append(cards, card)
		}
		if len(cards) >= 30 {
			break
		}
	}
	if len(cards) < 30 {
		t.Errorf("Random card should return different cards, got %v different", len(cards))
	}
}
func TestCardFromString(t *testing.T) {
	_, err := cardFromString("0 0")
	if (err != nil) {
		t.Errorf("Card from string 0 0 should not error");
	}
}