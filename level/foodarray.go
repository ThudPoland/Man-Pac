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
func (array *FoodArray) GenerateFoodArray(level *Level, characters []basic.Character) {
	for element := range level.layout {
		for nestedElement := range level.layout[element] {
			if level.layout[element][nestedElement] == 0 {
				var food basic.Food
				food.X = nestedElement
				food.Y = element
				array.array = append(array.array, food)
			}
		}
	}
}

//Draw implements drawing of food
func (array *FoodArray) Draw(target pixel.Target, manager sprite.Manager) {
	for element := range array.array {
		foodIndex := manager.GetSpriteIndexByName("food")
		destination := pixel.IM.Moved(pixel.V(16.0, 16.0)).
			Moved(pixel.V(float64(array.array[element].X*manager.GetSpriteSize()),
				float64(array.array[element].Y*manager.GetSpriteSize())))
		manager.DrawSprite(foodIndex, target, destination)
	}
}
