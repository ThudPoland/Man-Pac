package game

import (
	"github.com/ThudPoland/Man-Pac/basic"
)

//PreviousCharacter sets previous character
func PreviousCharacter(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.resources.DecrementIndex()
	}
}

//NextCharacter sets next character
func NextCharacter(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.resources.IncrementIndex()
	}
}

//SetUpDirection sets up direction for actual ghost
func SetUpDirection(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.SetDirection(basic.Up)
	}
}

//SetDownDirection sets down direction for actual ghost
func SetDownDirection(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.SetDirection(basic.Down)
	}
}

//SetLeftDirection sets left direction for actual ghost
func SetLeftDirection(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.SetDirection(basic.Left)
	}
}

//SetRightDirection sets right direction for actual ghost
func SetRightDirection(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.SetDirection(basic.Right)
	}
}

//SetNoDirection sets right direction for actual ghost
func SetNoDirection(receiver interface{}) {
	gameProvider, ok := receiver.(*Game)
	if ok == true {
		gameProvider.SetDirection(basic.No)
	}
}
