package basic

//Direction defines direction for actor
type Direction int

const (
	//No points no direction
	No Direction = iota
	//Up points the up direction
	Up = iota
	//Left points the left direction
	Left = iota
	//Down points the down direction
	Down = iota
	//Right points the right direction
	Right = iota
)

//SetDirection sets direction of entity based on actual position
func (direction *Direction) SetDirection(points NearPoints, inputDirection Direction) {
	*direction = No
	switch inputDirection {
	case Up:
		if points.Up == 1 {
			return
		}
	case Down:
		if points.Down == 1 {
			return
		}
	case Right:
		if points.Right == 1 {
			return
		}
	case Left:
		if points.Left == 1 {
			return
		}
	}
	*direction = inputDirection
}
