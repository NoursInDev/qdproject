package noise

import (
	"testing"
)

func TestPerlin2DMap(t *testing.T) {
	width, height, numBiomes := 5, 5, 4
	result := Perlin2DMap(width, height, numBiomes)
	// verif largeur
	if len(result) != height {
		t.Errorf("La hauteur de la carte est incorrecte. Attendu: %d, Obtenu: %d", height, len(result))
	}

	// verif longueur
	for i, row := range result {
		if len(row) != width {
			t.Errorf("La longueur de la ligne %d est incorrecte. Attendu: %d, Obtenu: %d", i, width, len(row))
		}
	}

	// verif plage biome
	for i, row := range result {
		for j, biome := range row {
			if biome < 0 || biome >= numBiomes {
				t.Errorf("La valeur de biome à la position (%d, %d) est hors de la plage attendue: %d", i, j, biome)
			}
		}
	}
}
