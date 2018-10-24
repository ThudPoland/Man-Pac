package basic

//Actor contains methods for
type Actor interface {
	GetPosition() (int, int)
	GetX() int
	GetY() int

	GetSize() (int, int)
	GetWidth() int
	GetHeight() int

	Move(x int, y int)
	SetPosition(x int, y int)
	SetSize(w int, h int)
}
