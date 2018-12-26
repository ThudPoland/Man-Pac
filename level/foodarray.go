package level

import (
	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//FoodArray implements array for food
type FoodArray struct {
	array []basic.Food
}

//GenerateFoodArray generates array of food.
func (array *FoodArray) GenerateFoodArray(level *Level, characters []basic.DrawableActor) {
	array.array = nil
	for element := range level.layout {
		for nestedElement := range level.layout[element] {
			if level.layout[element][nestedElement] == 0 {
				addElement := true
				var food basic.Food
				food.X = nestedElement
				food.Y = element

				nearPoints := level.GetNearPoints(nestedElement, element)

				food.EstimateFoodType(nearPoints)

				for characterIndex := range characters {
					characterX := characters[characterIndex].GetX()
					characterY := characters[characterIndex].GetY()
					if food.X == characterX && food.Y == characterY {
						addElement = false
						break
					}
				}
				if addElement {
					array.array = append(array.array, food)
				}
			}
		}
	}
}

//Draw implements drawing of food
func (array *FoodArray) Draw(target pixel.Target, manager sprite.Manager) {
	for element := range array.array {
		var foodIndex int
		if array.array[element].GetFoodType() == basic.PowerUp {
			foodIndex = manager.GetSpriteIndexByName("powerup")
		} else {
			foodIndex = manager.GetSpriteIndexByName("food")
		}
		destination := pixel.IM.Moved(pixel.V(16.0, 16.0)).
			Moved(pixel.V(float64(array.array[element].X*manager.GetSpriteSize()),
				float64(array.array[element].Y*manager.GetSpriteSize())))
		manager.DrawSprite(foodIndex, target, destination)
	}
}

//Eat is function for eating food
func (array *FoodArray) Eat(x, y int) {
	if array != nil {
		for element := range array.array {
			if array.array[element].X == x && array.array[element].Y == y {
				array.array = append(array.array[:element], array.array[element+1:]...)
				break
			}
		}
	}
}

//GetDirectionForFood gets direction for food next to point
func (array *FoodArray) GetDirectionForFood(x, y int) basic.Direction {
	possibleDirections := []basic.Direction{basic.Left, basic.Down, basic.Up, basic.Right}
	directionChanges := []struct{ x, y int }{{x - 1, y}, {x, y - 1}, {x, y + 1}, {x + 1, y}}
	for element := range possibleDirections {
		for nestedElement := range array.array {
			if array.array[nestedElement].X == directionChanges[element].x && array.array[nestedElement].Y == directionChanges[element].y {
				return possibleDirections[element]
			}
		}
	}

	return basic.No
}
