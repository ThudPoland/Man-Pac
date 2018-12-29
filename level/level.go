//Package level provides structure for creating levels
package level

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Level struct is struct for creating levels
type Level struct {
	name      string
	layout    [][]int
	render    pixel.Batch
	offset    pixel.Vec
	foodArray FoodArray
}

func (level *Level) parse(data Data) {
	for element := range data.lines {
		level.parseLine(data.lines[element])
	}
}

func (level *Level) parseLine(line string) {
	data := strings.Map(func(element rune) rune {
		if unicode.IsNumber(element) || element == ',' {
			return element
		}
		return -1
	}, line)

	splittedData := strings.Split(data, ",")
	levelData := make([]int, len(splittedData))
	for element := range splittedData {
		levelData[element], _ = strconv.Atoi(splittedData[element])
	}

	level.layout = append(level.layout, levelData)
}

//Load function for loading level
func (level *Level) Load(path string) {
	var data Data
	data.Load(path)
	level.parse(data)
}

//SetOffset sets offset
func (level *Level) SetOffset(offset pixel.Vec) {
	level.offset = offset
}

//Draw method for level
func (level Level) Draw(target pixel.Target, manager sprite.Manager) {
	for x := range level.layout {
		for y := range level.layout[x] {
			destination := pixel.IM.Moved(pixel.V(float64(y*manager.GetSpriteSize())+16,
				float64(x*manager.GetSpriteSize())+16)).Moved(level.offset)
			manager.DrawSprite(level.layout[x][y], target, destination)
		}
	}
	level.foodArray.Draw(target, manager)
}

//GetNearPoints gets near points
func (level *Level) GetNearPoints(x, y int) basic.NearPoints {
	return basic.NearPoints{Position: pixel.V(float64(x), float64(y)),
		Up:    level.layout[y+1][x],
		Down:  level.layout[y-1][x],
		Left:  level.layout[y][x-1],
		Right: level.layout[y][x+1]}
}

//GenerateFoodArray generates food array for the level
func (level *Level) GenerateFoodArray(characters []basic.DrawableActor) {
	level.foodArray.GenerateFoodArray(level, characters)
}

//Eat is function from eating things from a level
func (level *Level) Eat(x int, y int) {
	level.foodArray.Eat(x, y)
}

//GetDirectionForFood gets direction for food next to point
func (level *Level) GetDirectionForFood(x, y int) basic.Direction {
	return level.foodArray.GetDirectionForFood(x, y)
}

//GetLayout gets layout of a level
func (level *Level) GetLayout() *[][]int {
	return &level.layout
}

//GetAnyFoodData gets any food
func (level *Level) GetAnyFoodData() *basic.Food {
	return level.foodArray.GetAnyFoodData()
}
