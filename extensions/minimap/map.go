package minimap

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"

	_ "embed"
)

//go:embed mapids.png
var mapBytes []byte

var MapImage *ebiten.Image

func Load() {
	decoded, _, err := image.Decode(bytes.NewReader(mapBytes))
	if err != nil {
		log.Fatal(err)
	}
	MapImage = ebiten.NewImageFromImage(decoded)
}
