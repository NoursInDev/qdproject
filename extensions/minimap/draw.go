package minimap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"qdproject/extensions/global"
)

func Draw(screen *ebiten.Image, fmapcontent [][]int, xchar, ychar int) {
	for x := range fmapcontent {
		for y := range fmapcontent[x] {
			if fmapcontent[x][y] != -1 {
				op := &ebiten.DrawImageOptions{}                                                          // nouvelle instance DrawImageOptions
				op.GeoM.Translate(float64(y*global.Data.MapTileSize), float64(x*global.Data.MapTileSize)) // matrice de géométries
				op.GeoM.Scale(
					global.Data.MapDimX/(float64(len(fmapcontent[0]))*float64(global.Data.MapTileSize)),
					global.Data.MapDimY/(float64(len(fmapcontent))*float64(global.Data.MapTileSize)),
				)
				// Utilisation de x*configuration.Global.TileSize pour le décalage en x (pour afficher le bon bloc)
				var shiftX int
				if y == xchar && x == ychar {
					shiftX = 4 * global.Data.MapTileSize
				} else {
					shiftX = fmapcontent[x][y] * global.Data.MapTileSize
				} // décalage en x sur l'image
				img := MapImage.SubImage(image.Rect(shiftX, 0, shiftX+global.Data.MapTileSize, global.Data.MapTileSize)).(*ebiten.Image) //récupération de la sous image
				screen.DrawImage(img, op)
			}
		}
	}
}
