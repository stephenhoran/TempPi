package main

import (
	"TemperaturePi/fonts"
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	hemicubeFont font.Face
)

//nolint:gochecknoinits
func init() {
	tt, err := opentype.Parse(fonts.HemicubeFont)
	if err != nil {
		log.Fatalln(err)
	}

	hemicubeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{Size: 24, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		log.Fatalln(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, "hello world", hemicubeFont, 20, 40, color.White)
}

func (g *Game) Layout(width, height int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Temperature Pi")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalln(err)
	}
}
