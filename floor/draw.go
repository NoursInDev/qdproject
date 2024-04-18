package floor

import (
	"image"

	"qdproject/assets"
	"qdproject/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image) {
	for y := range f.content {
		for x := range f.content[y] {
			if f.content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

				var img *ebiten.Image
				if f.content[y][x] > 76 && f.content[y][x] < 88 {
					shiftX := (f.content[y][x] - 77) * configuration.Global.TileSize
					img = assets.TpImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image)
				} else if f.content[y][x] > 4 && f.content[y][x] < 11 {
					shiftX := (f.content[y][x] - 5) * configuration.Global.TileSize
					img = assets.WaterImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image)

				} else {
					shiftX := f.content[y][x] * configuration.Global.TileSize
					img = assets.FloorImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image)
				}

				screen.DrawImage(img, op)

			}
		}
	}
}