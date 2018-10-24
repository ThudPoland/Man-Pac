package game

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Executer provides all functions for execution of game
type Executer interface {
	SetActualLevel(index int)
	ProcessStateOfActors()
	Draw(t pixel.Target)
	SetLevelManager(spriteManager *sprite.Manager)
}
