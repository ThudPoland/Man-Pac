package ai

import (
	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/pathfinding"
)

//FoodFinder makes DumbEater clever with pathfinding algorithms
type FoodFinder struct {
	DumbEater
}

//Calculate calculates next direction for AI
func (behaviour *FoodFinder) Calculate(character *basic.Character) {
	behaviour.direction = basic.No
	pathfindinder := &pathfinding.BreadthFirstSearch{}
	levelGraph := pathfinding.InitGraph()
	levelGraph.FillGraph(*behaviour.level.GetLayout())
	foodToEat := behaviour.level.GetAnyFoodData()
	if foodToEat != nil {
		pathfindinder.GetSearchResult(character.GetX(), character.GetY(), foodToEat.X, foodToEat.Y, levelGraph)
	}
	behaviour.direction = pathfindinder.GetDirection()
	if behaviour.direction == basic.No {
		behaviour.Dumb.Calculate(character)
	}
}
