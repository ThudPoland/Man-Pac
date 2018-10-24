package basic

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Drawable implements method for draw
type Drawable interface {
	Draw(t pixel.Target, destination pixel.Matrix, manager sprite.Manager)
}
