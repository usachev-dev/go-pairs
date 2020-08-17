package main

import (
	"../../game"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	textFont := getFont("/ArialMT.ttf", game.ScoreFontSize)
	cardFont := getFont("/font.ttf", game.ScoreFontSize*3)
	g := game.NewGame(textFont, cardFont)
	g.StartGame()
}

func getFont(url string, fontSize int) font.Face {
	textFontFileResponse, getErr := http.Get(url)
	if getErr != nil {
		panic("Could not get font file")
	}
	//var fontData []byte;
	fontData, readErr := ioutil.ReadAll(textFontFileResponse.Body)
	if readErr != nil {
		panic("Could not read response")
	}
	tt, parseErr := truetype.Parse(fontData)
	if parseErr != nil {
		panic("Could not parse font format")
	}
	return truetype.NewFace(tt, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     300,
		Hinting: font.HintingNone,
	})
}
