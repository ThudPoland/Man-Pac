package sprite

import (
	"github.com/faiface/pixel"
)

type Drawer interface {
	Draw(target pixel.Target, destination pixel.Matrix)
}

//Draw function draws sprite
func (sprite Sprite) Draw(target pixel.Target, destination pixel.Matrix) {
	if sprite.sprite != nil {
		sprite.sprite.Draw(target, destination)
	}
}
