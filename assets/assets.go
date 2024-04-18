package assets

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed floor.png
var floorBytes []byte

//go:embed water.png
var waterBytes []byte

//go:embed tpanim2.png
var tpBytes []byte

// FloorImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents éléments qui peuvent s'afficher au sol
// (herbe, sable, etc).
// Dans la version du projet qui vous est fournie, ces éléments sont des
// carrés de 16 pixels de côté. Vous pourrez changer cela si vous le voulez.

// Ajout des textures de l'eau sous 5 formes à la suite sous water.png
var FloorImage *ebiten.Image
var WaterImage *ebiten.Image
var TpImage *ebiten.Image
var ParticuleImage *ebiten.Image

//go:embed characters/character_alt.png
var characterBytes []byte

// CharacterImage contient une version compatible avec Ebitengine de
// l'image qui contient les différentes étapes de l'animation du
// personnage.
// Dans la version du projet qui vous est fournie, ce personnage tient
// dans un carré de 16 pixels de côté. Vous pourrez changer cela si vous
// le voulez.
var CharacterImage *ebiten.Image

// Load est la fonction en charge de transformer, à l'exécution du programme,
// les images du jeu en structures de données compatibles avec Ebitengine.
// Ces structures de données sont stockées dans les variables définies ci-dessus.
func Load() {
	decoded, _, err := image.Decode(bytes.NewReader(floorBytes))
	if err != nil {
		log.Fatal(err)
	}
	FloorImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(waterBytes))
	if err != nil {
		log.Fatal(err)
	}
	WaterImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(tpBytes))
	if err != nil {
		log.Fatal(err)
	}
	TpImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(characterBytes))
	if err != nil {
		log.Fatal(err)
	}
	CharacterImage = ebiten.NewImageFromImage(decoded)
}
