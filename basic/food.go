package basic

//FoodType defines enum for food types
type FoodType int

const (
	//CommonMeal is enum for normal type of food
	CommonMeal FoodType = iota
	//PowerUp is food which is also power-up
	PowerUp = iota
	//RichMeal is food with many points to earn
	RichMeal = iota
)

//Food defines food
type Food struct {
	PhysicalObject
	foodType FoodType
}

//SetFoodType sets food type
func (food *Food) SetFoodType(foodType FoodType) {
	food.foodType = foodType
}

//GetFoodType gets food type
func (food *Food) GetFoodType() FoodType {
	return food.foodType
}

//EstimateFoodType estimates type of food
func (food *Food) EstimateFoodType(points NearPoints) {
	if points.Sum() == 3 {
		food.SetFoodType(PowerUp)
	} else {
		food.SetFoodType(CommonMeal)
	}
}
