package ai

import (
	"math/rand"

	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/level"
)

//Dumb is structure implementing dumb (random-driven) AI
type Dumb struct {
	level     *level.Level
	valid     bool
	direction basic.Direction
}

//SetInput sets input for dumb algorithm
func (behaviour *Dumb) SetInput(data interface{}) {
	level, ok := data.(*level.Level)
	if !ok && level != nil {
		behaviour.valid = false
	}
	behaviour.level = level
	behaviour.valid = true

}

//Calculate calculates next direction for AI
func (behaviour *Dumb) Calculate(character *basic.Character) {
	possibleDirections := []basic.Direction{basic.Left, basic.Down, basic.Up, basic.Right}
	points := behaviour.level.GetNearPoints(character.GetX(), character.GetY())
	behaviour.direction.SetDirection(points, behaviour.direction)
	if points.Sum() >= 4 {
		return
	}
	if behaviour.direction == basic.No || points.Sum() < 2 {
		for {
			value := rand.Intn(4)
			behaviour.direction.SetDirection(points, possibleDirections[value])
			if behaviour.direction != basic.No {
				break
			}
		}
	}
}

//GetDirection gets next direction for AI
func (behaviour *Dumb) GetDirection() basic.Direction {
	return behaviour.direction
}
