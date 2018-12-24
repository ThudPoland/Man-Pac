package basic

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//ManPac is enemy actor
type ManPac struct {
	Character
	spriteManager *sprite.Manager
}

//Draw draws character
func (enemy *ManPac) Draw(t pixel.Target, offset pixel.Vec, manager *sprite.Manager) {
	enemy.Character.Draw(t, offset, manager)
}

//SetSpriteManager sets sprite manager for ManPac
func (enemy *ManPac) SetSpriteManager(spriteManager *sprite.Manager) {
	enemy.spriteManager = spriteManager
}

//ProcessTurn is used to processing turn
func (enemy *ManPac) ProcessTurn() {
	switch enemy.direction {
	case Up:
		enemy.Y++
	case Down:
		enemy.Y--
	case Left:
		enemy.X--
	case Right:
		enemy.X++
	}
}
