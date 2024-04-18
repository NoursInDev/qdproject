package minimap

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestDraw(t *testing.T) {
	const (
		mapTileSize = 32
		mapDimX     = 800
		mapDimY     = 600
	)

	testImage := ebiten.NewImage(mapDimX, mapDimY)

	testImage.Fill(color.RGBA{255, 255, 255, 255})

	testMap := [][]int{
		{0, 1, 2, 3},
		{3, 2, 1, 0},
		{0, 1, 2, 3},
		{3, 2, 1, 0},
	}

	Draw(testImage, testMap)

	if c := testImage.At(0, 0); c == (color.RGBA{255, 255, 255, 255}) {
		t.Errorf("Attendu non blanc > Obtenu blanc")
	}
}

// note: crash à l'appel (not fixed)
