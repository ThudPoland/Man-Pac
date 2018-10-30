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
