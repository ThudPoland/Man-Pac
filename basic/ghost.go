package basic

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Ghost is user-controlled actor
type Ghost struct {
	Character
	direction        Direction
	interfaceManager *sprite.Manager
}

//Draw draws character
func (ghost *Ghost) Draw(t pixel.Target, offset pixel.Vec, manager *sprite.Manager) {
	ghost.Character.Draw(t, offset, manager)
	destination := pixel.IM.Moved(offset).Moved(pixel.V(float64(ghost.X*manager.GetSpriteSize()),
		float64(ghost.Y*manager.GetSpriteSize())))
	if ghost.interfaceManager != nil {
		switch ghost.direction {
		case No:
		case Up:
			ghost.interfaceManager.DrawSpriteByName("up", t, destination)
		case Down:
			ghost.interfaceManager.DrawSpriteByName("down", t, destination)
		case Left:
			ghost.interfaceManager.DrawSpriteByName("left", t, destination)
		case Right:
			ghost.interfaceManager.DrawSpriteByName("right", t, destination)
		default:
		}
	}
}

//SetDirection sets ghost direction
func (ghost *Ghost) SetDirection(direction Direction, points NearPoints) {
	switch direction {
	case Up:
		if points.Up == 1 {
			return
		}
	case Down:
		if points.Down == 1 {
			return
		}
	case Right:
		if points.Right == 1 {
			return
		}
	case Left:
		if points.Left == 1 {
			return
		}
	}
	ghost.direction = direction
}

//SetNoDirection sets no direction
func (ghost *Ghost) SetNoDirection() {
	ghost.direction = No
}

//SetInterfaceManager sets interface manager
func (ghost *Ghost) SetInterfaceManager(interfaceManager *sprite.Manager) {
	ghost.interfaceManager = interfaceManager
}
