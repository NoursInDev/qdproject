package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"qdproject/configuration"
	"qdproject/game"

	"qdproject/assets"

	"qdproject/extensions/global"
	"qdproject/extensions/minimap"
)

func main() {

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")

	flag.Parse()

	configuration.Load(configFileName)
	global.Load(configFileName)

	assets.Load()

	minimap.Load()

	g := &game.Game{}
	log.Println("game.Game éxécuté.")
	g.Init()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Let Us Cook Now - QDPROJECT")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
