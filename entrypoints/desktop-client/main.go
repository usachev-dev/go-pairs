package main

import (
	"../../game"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"time"
)

func loadFont(filename string,  fontSize int) font.Face {
	fontFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read font")
	}
	tt, err := truetype.Parse(fontFile)
	if err != nil {
		panic("could not parse font")
	}
	return truetype.NewFace(tt, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     300,
		Hinting: font.HintingNone,
	})
}

func main() {
	rand.Seed(time.Now().Unix())
	cardFont := loadFont("./assets/font.ttf", game.CardFontSize)
	textFont := loadFont("./assets/ArialMT.ttf", game.ScoreFontSize)
	g := game.NewGame(textFont, cardFont)
	g.StartGame()
}
