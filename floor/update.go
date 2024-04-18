package floor

import (
	"log"
	"qdproject/configuration"
	"qdproject/extensions/global"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor, randomFloor, biomeFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	//chargement
	if f.FullContent == nil {
		log.Println("Erreur : FullContent non chargé depuis le fichier.")
		return
	}

	//taille du content
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	// f.content remplissage
	for y := 0; y < configuration.Global.NumTileY; y++ {
		for x := 0; x < configuration.Global.NumTileX; x++ {
			absX := camXPos + x - configuration.Global.ScreenCenterTileX
			absY := camYPos + y - configuration.Global.ScreenCenterTileY
			if absX >= 0 && absX < len(f.FullContent[0]) && absY >= 0 && absY < len(f.FullContent) {
				//copie valeurs
				f.content[y][x] = f.FullContent[absY][absX]
				if f.content[y][x] > 4 && f.content[y][x] < 11 { //traitement eau
					if global.Data.Tick%12 == 0 {
						f.FullContent[absY][absX]++
						if f.FullContent[absY][absX] == 10 {
							f.FullContent[absY][absX] = 5
						}
					}
				} else if f.content[y][x] > 76 && f.content[y][x] < 88 {
					if global.Data.Tick%6 == 0 {
						f.FullContent[absY][absX]++
						if f.FullContent[absY][absX] == 87 {
							f.FullContent[absY][absX] = 77
						}
					}
				}
			} else {
				//case vide
				f.content[y][x] = -1
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
