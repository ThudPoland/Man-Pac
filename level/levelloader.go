package level

import (
	"strconv"
	"strings"
	"unicode"
)

//Loader for level interface
type Loader interface {
	Load(path string)
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
func (level Level) Load(path string) {
	var data Data
	data.Load(path)
	level.parse(data)
}
