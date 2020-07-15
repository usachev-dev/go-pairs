package main

import (
	"../../game"
	"errors"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth   = 640
	screenHeight  = 640
	cardFontSize  = 30
	scoreFontSize = 10
	cardHeight    = cardFontSize * 3.2
	cardWidth     = cardFontSize * 2.4
	boardWidth    = cardWidth * 6
	boardHeight   = cardHeight * 6
)

var fontFace font.Face
var arial font.Face

type Game struct {
	deck                game.Deck
	pairRevealCountDown int
	revealedCardIndexes []int
	scoredCardIndexes   []int
	tries               int
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
		measure:= text.MeasureString(winText, arial)
		w, h := screen.Size()
		x := w/2 - measure.X/2
		y := h/2 - measure.Y/2
		text.Draw(screen, winText, arial, x, y, color.Black)
		return
	}

	for i, card := range g.deck {
		if cardIsScored(i, g) {
			continue
		}
		if cardIsRevealed(i, g) {
			drawCard(card, screen, i)
		} else {
			//drawCard(card, screen, i)
			drawCardBack(card, screen, i)
		}
	}

	score := len(g.scoredCardIndexes) / 2
	scoreStr := fmt.Sprintf("Tries: %d, Score: %d", g.tries, score)
	text.Draw(screen, scoreStr, arial, 0, scoreFontSize*3, color.Black)
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

func drawCard(card game.Card, screen *ebiten.Image, index int) {
	screenWidth, screenHeight := screen.Size()
	x := cardX(screenWidth, index)
	y := cardY(screenHeight, index)
	text.Draw(screen, card.Unicode(), fontFace, x, y, card.Color())
}

func drawCardBack(card game.Card, screen *ebiten.Image, index int) {
	screenWidth, screenHeight := screen.Size()
	x := cardX(screenWidth, index)
	y := cardY(screenHeight, index)
	text.Draw(screen, card.CardBack(), fontFace, x, y, card.BackColor())
}

func loadFont(filename string, ff *font.Face, fontSize int) {
	fontFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read font")
	}
	tt, err := truetype.Parse(fontFile)
	if err != nil {
		panic("could not parse font")
	}
	*ff = truetype.NewFace(tt, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     300,
		Hinting: font.HintingNone,
	})
}

func main() {
	rand.Seed(time.Now().Unix())
	//filename, err := findfont.Find("micross.ttf")
	//if err !=nil {
	//	panic("Could not find font")
	//}
	loadFont("./assets/font.ttf", &fontFace, cardFontSize)
	loadFont("./assets/ArialMT.ttf", &arial, scoreFontSize)
	g := Game{
		game.NewDeckPairs()/*.Shuffle()*/,
		0,
		[]int{},
		[]int{},
		0,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pairs")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}

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
