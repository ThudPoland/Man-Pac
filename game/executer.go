package game

import (
	"github.com/faiface/pixel"
)

//Executer provides all functions for execution of game
type Executer interface {
	ProcessStateOfActors()
	Draw(t pixel.Target)
}
