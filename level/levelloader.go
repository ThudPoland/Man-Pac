package level

import (
	"strings"
	"unicode"
)

//Loader for level interface
type Loader interface {
	Parse(data Data)
	ParseLine(line string)
	Load(path string)
}

/*//Parse data for level
func (level Level) Parse(data Data) {
	for element in
}*/

func (level Level) ParseLine(line string) {
	strings.Map(func(element rune) rune {
		if unicode.IsNumber(element) || element == ',' {
			return element
		}
		return -1
	}, line)
}
