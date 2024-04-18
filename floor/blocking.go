package floor

import (
	"qdproject/configuration"
)

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	if configuration.Global.FloorKind != quadTreeFloor {
		blocking[0] = characterYPos <= 0 || f.FullContent[characterYPos-1][characterXPos] == -1 || (f.FullContent[characterYPos-1][characterXPos] > 4 && f.FullContent[characterYPos-1][characterXPos] < 11)
		blocking[1] = characterXPos >= len(f.FullContent[0])-1 || f.FullContent[characterYPos][characterXPos+1] == -1 || (f.FullContent[characterYPos][characterXPos+1] > 4 && f.FullContent[characterYPos][characterXPos+1] < 11)
		blocking[2] = characterYPos >= len(f.FullContent)-1 || f.FullContent[characterYPos+1][characterXPos] == -1 || (f.FullContent[characterYPos+1][characterXPos] > 4 && f.FullContent[characterYPos+1][characterXPos] < 11)
		blocking[3] = characterXPos <= 0 || f.FullContent[characterYPos][characterXPos-1] == -1 || (f.FullContent[characterYPos][characterXPos-1] > 4 && f.FullContent[characterYPos][characterXPos-1] < 11)
	} else {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}
	return blocking
}
