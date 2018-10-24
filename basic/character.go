package basic

//Character is character controlled used in gameplay
type Character struct {
	PhysicalObject
}

//GetPosition gets position of character
func (character Character) GetPosition() (int, int) {
	return character.X, character.Y
}

//GetX gets x position of character
func (character Character) GetX() int {
	return character.X
}

//GetY gets y position of character
func (character Character) GetY() int {
	return character.Y
}

//GetSize gets size of character
func (character Character) GetSize() (int, int) {
	return character.W, character.H
}

//GetWidth gets width of character
func (character Character) GetWidth() int {
	return character.W
}

//GetHeight gets height of character
func (character Character) GetHeight() int {
	return character.H
}

//Move moves a character
func (character *Character) Move(x int, y int) {
	character.X += x
	character.Y += y
}

//SetPosition sets position of character
func (character *Character) SetPosition(x int, y int) {
	character.X = x
	character.Y = y
}

//SetSize sets size of character
func (character *Character) SetSize(w int, h int) {
	character.W = w
	character.H = h
}
