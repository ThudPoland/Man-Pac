package level

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Drawer provides draw methods for level
type Drawer interface {
	Draw(batch pixel.Batch)
}

//Draw method for level
func (level Level) Draw(batch pixel.Batch, manager sprite.Manager) {

}
