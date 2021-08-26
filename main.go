package main

import (
	"TemperaturePi/fonts"
	"TemperaturePi/weather"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	WindowSizeX = 640
	WindowSizeY = 480
)

var (
	maiaFont    font.Face
	weatherData *weather.Weather
)

//nolint:gochecknoinits
func init() {
	tt, err := opentype.Parse(fonts.Maia_ttf)
	if err != nil {
		log.Fatalln(err)
	}

	maiaFont, err = opentype.NewFace(tt, &opentype.FaceOptions{Size: 24, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	var err error

	weatherData, err = weather.NewWeather()
	if err != nil {
		log.Println(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, "current: "+weatherData.Temperature.String(), maiaFont, 20, 40, color.White)
	text.Draw(screen, "Feels Like: "+weatherData.FeelsLike.String(), maiaFont, 20, 80, color.White)
	text.Draw(screen, "Min: "+weatherData.TemperatureMin.String(), maiaFont, 20, 120, color.White)
	text.Draw(screen, "Max: "+weatherData.TemperatureMax.String(), maiaFont, 20, 160, color.White)
}

func (g *Game) Layout(width, height int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(WindowSizeX, WindowSizeY)
	ebiten.SetWindowTitle("Temperature Pi")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalln(err)
	}

}
