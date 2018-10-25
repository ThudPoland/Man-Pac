package level

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Drawer provides draw methods for level
type Drawer interface {
	Draw(target pixel.Target, manager sprite.Manager)
	SetOffset(offset pixel.Vec)
}
