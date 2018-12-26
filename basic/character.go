package basic

import (
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Character is character controlled used in gameplay
type Character struct {
	PhysicalObject
	direction   Direction
	spriteIndex int
}

//GetPosition gets position of character
func (character Character) GetPosition() (int, int) {
	return character.X, character.Y
}

//GetX gets x position of character
func (character *Character) GetX() int {
	return character.X
}

//GetY gets y position of character
func (character *Character) GetY() int {
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

//Draw draws character
func (character *Character) Draw(t pixel.Target, offset pixel.Vec, manager *sprite.Manager) {
	destination := pixel.IM.Moved(offset).Moved(pixel.V(float64(character.X*manager.GetSpriteSize()),
		float64(character.Y*manager.GetSpriteSize())))
	manager.DrawSprite(character.spriteIndex, t, destination)
}

//SetIndexForSprite sets index for sprite from sprites manager
func (character *Character) SetIndexForSprite(name string, manager *sprite.Manager) {
	spriteIndex := manager.GetSpriteIndexByName(name)
	if spriteIndex != -1 {
		character.spriteIndex = spriteIndex
	}
}

//IsValid checks if character is valid for game
func (character *Character) IsValid() bool {
	if character.spriteIndex != -1 {
		return true
	}
	return false
}

//SetDirection sets character direction
func (character *Character) SetDirection(direction Direction, points NearPoints) {
	character.direction.SetDirection(points, direction)
}

//SetNoDirection sets no direction
func (character *Character) SetNoDirection() {
	character.direction = No
}

//GetDirection gets direction
func (character *Character) GetDirection() Direction {
	return character.direction
}

//ProcessTurn processes turn for character
func (character *Character) ProcessTurn() {
	switch character.direction {
	case Up:
		character.Y++
	case Down:
		character.Y--
	case Left:
		character.X--
	case Right:
		character.X++
	}
	character.direction = No
}

//IsReady is function for checking if character is ready to process
func (character *Character) IsReady() bool {
	return character.direction != No
}
