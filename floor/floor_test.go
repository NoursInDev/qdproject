package floor

import (
	"log"
	"testing"
)

func TestReadFloorFromFile(t *testing.T) {
	got := readFloorFromFile("../floor-files/test")
	want := [][]int{
		{1, 1, 0, 1}, {2, 0, 0, 1}, {3, 7, 1, 5}, {1, 1, 2, 0},
	}

	if !isEqual(got, want) {
		t.Errorf("Failed: renvoyé -> %d, attendue -> %d", got, want)
	} else {
		log.Printf("ReadFloorFromFile: Success")
	}
}

func TestFloorGeneration(t *testing.T) {
	x := 10
	y := 20
	got := floorGeneration(y, x)
	if len(got) != y {
		t.Errorf("Failed: taille totale %d, attendue -> %d", len(got), y)
	} else {
		for i := 0; i < len(got); i++ {
			if len(got[i]) != x {
				t.Errorf("Failed: taille du tableau n°%d -> %d, attendue -> %d", i, len(got), y)
				break
			}
		}
		log.Printf("FloorGeneration: Success")
	}
}

func TestSaveFloorToFile(t *testing.T) {
	table := make([][]int, 0, 0)
	if SaveFloorToFile(table) != nil {
		t.Errorf("Failed: Erreur de sauvegarde.")
	}
	log.Printf("SaveFloorToFile: Success | ! Carte vierge sauvegardée.")
}

// comparaison tables
func isEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
