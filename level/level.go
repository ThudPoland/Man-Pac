//Package level provides structure for creating levels
package level

import (
	"github.com/faiface/pixel"
)

//Level struct is struct for creating levels
type Level struct {
	name   string
	layout [][]int
	render pixel.Batch
}
