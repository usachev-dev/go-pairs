package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
)

const (
	screenWidth   = 640
	screenHeight  = 640
	CardFontSize  = 30
	ScoreFontSize = 10
	cardHeight    = CardFontSize * 3.2
	cardWidth     = CardFontSize * 2.4
	boardWidth    = cardWidth * 6
	boardHeight   = cardHeight * 6
)

type Game struct {
	deck                Deck
	pairRevealCountDown int
	revealedCardIndexes []int
	scoredCardIndexes   []int
	tries               int
	textFont            font.Face
	cardFont            font.Face
}

func NewGame(textFont font.Face, cardFont font.Face) Game {
	return Game{
		NewDeckPairs().Shuffle(),
		0,
		[]int{},
		[]int{},
		0,
		textFont,
		cardFont,
	}
}

func (g Game) StartGame() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pairs")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.pairRevealCountDown > 0 {
		g.pairRevealCountDown--
		if g.pairRevealCountDown == 0 {
			if len(g.revealedCardIndexes) == 2 && g.deck[g.revealedCardIndexes[0]].Equals(g.deck[g.revealedCardIndexes[1]]) {
				g.scoredCardIndexes = append(g.scoredCardIndexes, g.revealedCardIndexes...)
			}
			g.revealedCardIndexes = []int{}
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if len(g.revealedCardIndexes) >= 2 {
			return nil
		}
		screenWidth, screenHeight := screen.Size()
		clickX, clickY := ebiten.CursorPosition()
		clickedCardIndex, err := clickCardIndex(clickX, clickY, screenWidth, screenHeight)
		if err != nil {
			return nil
		}
		if cardIsRevealed(clickedCardIndex, g) {
			return nil
		}
		g.revealedCardIndexes = append(g.revealedCardIndexes, clickedCardIndex)
		if len(g.revealedCardIndexes) >= 2 {
			g.tries++
			g.pairRevealCountDown = 30
		}
	}
	return nil
}

func cardIsRevealed(index int, game *Game) bool {
	for _, i := range game.revealedCardIndexes {
		if i == index {
			return true
		}
	}
	return false
}

func cardIsScored(index int, game *Game) bool {
	for _, i := range game.scoredCardIndexes {
		if i == index {
			return true
		}
	}
	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	//op := &ebiten.DrawImageOptions{}
	screen.Fill(color.RGBA{240, 240, 240, 255})

	if len(g.scoredCardIndexes) == len(g.deck) {
		winText := fmt.Sprintf("You won with %d tries!", g.tries)
		measure := text.MeasureString(winText, g.textFont)
		w, h := screen.Size()
		x := w/2 - measure.X/2
		y := h/2 - measure.Y/2
		text.Draw(screen, winText, g.textFont, x, y, color.Black)
		return
	}

	for i, card := range g.deck {
		if cardIsScored(i, g) {
			continue
		}
		if cardIsRevealed(i, g) {
			drawCard(card, screen, i, g.cardFont)
		} else {
			//drawCard(card, screen, i)
			drawCardBack(card, screen, i, g.cardFont)
		}
	}

	score := len(g.scoredCardIndexes) / 2
	scoreStr := fmt.Sprintf("Tries: %d, Score: %d", g.tries, score)
	text.Draw(screen, scoreStr, g.textFont, 0, ScoreFontSize*3, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func boardOffsetX(screenWidth int) int {
	return screenWidth/2 - boardWidth/2
}

func boardOffsetY(screenHeight int) int {
	return screenHeight/2 - boardHeight/2
}

func cardX(screenWidth int, cardIndex int) int {
	return boardOffsetX(screenWidth) + cardIndex%6*cardWidth
}

func cardY(screenHeight int, cardIndex int) int {
	return boardOffsetY(screenHeight) + cardHeight + (cardIndex/6)*cardHeight
}

func clickCardIndex(x int, y int, screenWidth int, screenHeight int) (int, error) {
	boardOffsetX := boardOffsetX(screenWidth)
	boardOffsetY := boardOffsetY(screenHeight)
	if x < boardOffsetX || y < boardOffsetY || x > boardOffsetX+boardWidth || y > boardOffsetY+boardHeight {
		return 0, errors.New("Click is outside board")
	}
	cardX := (x - boardOffsetX) / cardWidth
	cardY := (y - boardOffsetY) / cardHeight
	cardIndex := cardY*6 + cardX
	return cardIndex, nil
}

func drawCard(card Card, screen *ebiten.Image, index int, font font.Face) {
	screenWidth, screenHeight := screen.Size()
	x := cardX(screenWidth, index)
	y := cardY(screenHeight, index)
	text.Draw(screen, card.Unicode(), font, x, y, card.Color())
}

func drawCardBack(card Card, screen *ebiten.Image, index int, font font.Face) {
	screenWidth, screenHeight := screen.Size()
	x := cardX(screenWidth, index)
	y := cardY(screenHeight, index)
	text.Draw(screen, card.CardBack(), font, x, y, card.BackColor())
}

type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

// MouseStrokeSource is a StrokeSource implementation of mouse.
type MouseStrokeSource struct{}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

// TouchStrokeSource is a StrokeSource implementation of touch.
type TouchStrokeSource struct {
	ID int
}

func (t *TouchStrokeSource) Position() (int, int) {
	return ebiten.TouchPosition(t.ID)
}

func (t *TouchStrokeSource) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.ID)
}
