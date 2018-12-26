package ai

import (
	"github.com/ThudPoland/Man-Pac/basic"
)

//DumbEater is structure implementing dumb, but nice in eating - (random-driven) AI
type DumbEater struct {
	Dumb
}

//Calculate calculates next direction for AI
func (behaviour *DumbEater) Calculate(character *basic.Character) {
	lastDirection := behaviour.direction
	behaviour.direction = behaviour.level.GetDirectionForFood(character.GetX(), character.GetY())
	if behaviour.direction == basic.No {
		behaviour.direction = lastDirection
		behaviour.Dumb.Calculate(character)
	}
}
