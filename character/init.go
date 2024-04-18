package character

import (
	"qdproject/camera"
	"qdproject/configuration"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init() {
	c.animationStep = 1

	if configuration.Global.CameraMode == camera.Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	} /*else if configuration.Global.ForceCenterPlayer == true{
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}*/
}