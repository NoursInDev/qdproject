package noise

import (
	"math/rand"
	"qdproject/configuration"
	"qdproject/extensions/save"
)

// fonction rajoutée pour la génération de biomes.

func MapBiomeGeneration(x, y int) (fcontent, fmap [][]int) {
	perlinmap := Perlin2DMap(x, y, 4)
	m := make([][]int, len(perlinmap))
	for i := range m {
		m[i] = make([]int, len(perlinmap[0]))
	}

	for i := 0; i < len(perlinmap); i++ {
		for j := 0; j < len(perlinmap[0]); j++ {
			biomeid := perlinmap[i][j]
			switch {
			case biomeid == 0:
				rnumber := rand.Float64()
				if rnumber > 0.5 {
					m[i][j] = 3
				} else {
					m[i][j] = 4
				}
			case biomeid == 1:
				rnumber := rand.Float64()
				if rnumber > 0.9 {
					m[i][j] = 2
				} else {
					m[i][j] = 0
				}
			case biomeid == 2:
				m[i][j] = 1
			case biomeid == 3:
				m[i][j] = 5 + rand.Intn(4)
			case biomeid >= 4:
				m[i][j] = 5 + rand.Intn(4) // in comming, eau profonde
			}
		}
	}

	if configuration.Global.SaveMap == true {
		save.SaveFloorToFile(m)
	}
	return m, perlinmap

}
