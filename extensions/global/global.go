package global

import (
	"encoding/json"
	"log"
	"os"
)

type GlobalData struct {
	// extension animation eau (ticks)
	Tick int

	// extension minimap (liée aux biomes)
	Fmapcontent [][]int
	MapTileSize int
	MapDimX     float64
	MapDimY     float64
	Enabled     bool

	// taille gen map
	MapX int
	MapY int

	// teleporter
	Tp1x int
	Tp1y int
	Tp2x int
	Tp2y int

	Ptp1 int
	Ptp2 int

	Active bool
	NbTP   int

	TPNameExtension string
}

var Data GlobalData

func Load(configurationFileName string) {
	content, err := os.ReadFile(configurationFileName)
	if err != nil {
		log.Fatal("Error while opening configuration file: ", err)
	}

	err = json.Unmarshal(content, &Data)
	if err != nil {
		log.Fatal("Error while reading configuration file: ", err)
	}

}
