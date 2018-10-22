//Package sprite provides sprites for game
package sprite

import (
	"github.com/faiface/pixel"
)

//Sprite implements pixel's image as a part of game
type Sprite struct {
	picture     pixel.Picture
	sprite      *pixel.Sprite
	spriteError error
	id          int
}
