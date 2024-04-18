package quadtree

import (
	"qdproject/configuration"
	"qdproject/extensions/global"
)

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du quadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	//main part
	for i := 0; i < configuration.Global.NumTileX; i++ {
		for j := 0; j < configuration.Global.NumTileY; j++ {
			absX := topLeftX + i
			absY := topLeftY + j
			if absX < q.width && absY < q.height && absX >= 0 && absY >= 0 {
				contentHolder[j][i] = ReturnFloor(absX, absY, q.root)
			} else {
				contentHolder[j][i] = -1
			}
			if contentHolder[j][i] > 4 { //traitement eau
				if global.Data.Tick%12 == 0 {
					ChangeWaterValue(absX, absY, q.root)
				}
			}
		}
	}
}

func ReturnFloor(absX, absY int, n *node) int {
	//case feuille
	if n.topLeftNode == nil {
		return n.content
	}

	//determination coos carré suivant
	halfWidth := n.width / 2
	halfHeight := n.height / 2
	centerX := n.topLeftX + halfWidth
	centerY := n.topLeftY + halfHeight

	if absX < centerX {
		if absY < centerY {
			return ReturnFloor(absX, absY, n.topLeftNode)
		} else {
			return ReturnFloor(absX, absY, n.bottomLeftNode)
		}
	} else {
		if absY < centerY {
			return ReturnFloor(absX, absY, n.topRightNode)
		} else {
			return ReturnFloor(absX, absY, n.bottomRightNode)
		}
	}
}

func ChangeWaterValue(absX, absY int, n *node) {
	// case feuille
	if n.topLeftNode == nil {
		n.content++
		if n.content == 10 {
			n.content = 5
		}
		return
	}

	// détermination des coordonnées du carré suivant
	halfWidth := n.width / 2
	halfHeight := n.height / 2
	centerX := n.topLeftX + halfWidth
	centerY := n.topLeftY + halfHeight

	if absX < centerX {
		if absY < centerY {
			ChangeWaterValue(absX, absY, n.topLeftNode)
		} else {
			ChangeWaterValue(absX, absY, n.bottomLeftNode)
		}
	} else {
		if absY < centerY {
			ChangeWaterValue(absX, absY, n.topRightNode)
		} else {
			ChangeWaterValue(absX, absY, n.bottomRightNode)
		}
	}
}
