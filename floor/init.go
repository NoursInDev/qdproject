package floor

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"qdproject/configuration"
	"qdproject/extensions/global"
	"qdproject/extensions/noise"
	"qdproject/extensions/save"
	"qdproject/quadtree"
	"time"
)

func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case fromFileFloor:
		f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
		log.Printf("Carte chargée à partir du fichier: %s", configuration.Global.FloorFile)
	case quadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
		log.Printf("Carte chargée en utilisant un quadtree à partir du fichier: %s", configuration.Global.FloorFile)
	case randomFloor:
		f.FullContent = floorGeneration(global.Data.MapX, global.Data.MapY)
		log.Println("f.FullContent généré pour taille", len(f.FullContent), "x", len(f.FullContent[0]), "via génération aléatoire semi-simple")
	case biomeFloor:
		f.FullContent, global.Data.Fmapcontent = noise.MapBiomeGeneration(global.Data.MapX, global.Data.MapY)
		log.Println("f.FullContent généré pour taille", len(f.FullContent), "x", len(f.FullContent[0]), "via biome generation perlin noise")
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening floor file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, char := range line {
			row = append(row, int(char-'0'))
		}
		floorContent = append(floorContent, row)
	}
	return floorContent
}

func floorGeneration(NumTileY, NumTileX int) (floorContent [][]int) {
	// génération via seed
	rand.Seed(time.Now().UnixNano())

	// création carte vide
	floorContent = make([][]int, NumTileY)
	for i := range floorContent {
		floorContent[i] = make([]int, NumTileX)
	}

	// remplissage
	for y := 0; y < NumTileY; y++ {
		for x := 0; x < NumTileX; x++ {
			// gen aleatoire > détermination type de bloc
			randomValue := rand.Float64()
			// 90% de chances de prendre la valeur d'un bloc autour
			if randomValue <= 0.9 && y > 0 && x > 0 && y < NumTileY-1 && x < NumTileX-1 {
				neighbors := []int{
					floorContent[y-1][x-1], floorContent[y-1][x], floorContent[y-1][x+1],
					floorContent[y][x-1], floorContent[y][x+1],
					floorContent[y+1][x-1], floorContent[y+1][x], floorContent[y+1][x+1],
				}
				// Choisir aléatoirement parmi les blocs voisins
				floorContent[y][x] = neighbors[rand.Intn(len(neighbors))]
			} else {
				// 5% de chances de générer un chiffre entre 5 et 9 inclus
				if rand.Float64() <= 0.1 {
					floorContent[y][x] = rand.Intn(5) + 5 // Eau
				} else {
					// 95% de chances de générer un chiffre entre 0 et 4 inclus
					floorContent[y][x] = rand.Intn(5) // 0 à 4 (autres textures)
				}
			}
		}
	}

	// ne sauvegarder que si NumTileY et NumTileX sont supérieurs à 0
	if NumTileY > 0 && NumTileX > 0 && configuration.Global.SaveMap {
		save.SaveFloorToFile(floorContent)
	}
	return floorContent
}
