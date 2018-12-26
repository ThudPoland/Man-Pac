package basic

//Algorithm is interface for algorithm used by AI
type Algorithm interface {
	SetInput(data interface{})
	Calculate(character *Character)
	GetDirection() Direction
}
