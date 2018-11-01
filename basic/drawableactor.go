package basic

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//DrawableActor implements method for draw
type DrawableActor interface {
	Actor
	Processable
	Draw(t pixel.Target, offset pixel.Vec, manager *sprite.Manager)
}
