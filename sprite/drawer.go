package sprite

import (
	"github.com/faiface/pixel"
)

type Drawer interface {
	Draw(target pixel.Target, destination pixel.Matrix)
}
