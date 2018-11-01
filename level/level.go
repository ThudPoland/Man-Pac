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
	name   string
	layout [][]int
	render pixel.Batch
	offset pixel.Vec
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

//Main function for loading level
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
	for element := range level.layout {
		for nestedElement := range level.layout[element] {
			destination := pixel.IM.Moved(pixel.V(float64(nestedElement*manager.GetSpriteSize())+16,
				float64(element*manager.GetSpriteSize())+16)).Moved(level.offset)
			manager.DrawSprite(level.layout[element][nestedElement], target, destination)
		}
	}
}

//GetNearPoints gets near points
func (level *Level) GetNearPoints(x, y int) basic.NearPoints {
	return basic.NearPoints{Position: pixel.V(float64(x), float64(y)),
		Up:    level.layout[y+1][x],
		Down:  level.layout[y-1][x],
		Left:  level.layout[y][x-1],
		Right: level.layout[y][x+1]}
}
