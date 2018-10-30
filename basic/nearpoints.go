package basic

import (
	"github.com/faiface/pixel"
)

//NearPoints defines values of points near position
type NearPoints struct {
	Position pixel.Vec
	Up       int
	Down     int
	Left     int
	Right    int
}

//Sum sums all values of near points
func (points NearPoints) Sum() int {
	return points.Up + points.Down + points.Left + points.Right
}
