package level

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Drawer provides draw methods for level
type Drawer interface {
	Draw(target pixel.Target, manager sprite.Manager)
}

//Draw method for level
func (level Level) Draw(target pixel.Target, manager sprite.Manager) {
	for element := range level.layout {
		for nestedElement := range level.layout[element] {
			destination := pixel.IM.Moved(pixel.V(float64(nestedElement)*32+16, float64(element)*32+16))
			manager.DrawSprite(level.layout[element][nestedElement], target, destination)
		}
	}

}
