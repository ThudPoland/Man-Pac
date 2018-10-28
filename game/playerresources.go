package game

import (
	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//PlayerResources has info about all player resources
type PlayerResources struct {
	characters     []basic.Character
	manager        *sprite.Manager
	characterIndex int
}

//Draw draws player resources
func (resources PlayerResources) Draw(target pixel.Target) {
	for element := range resources.characters {
		resources.characters[element].Draw(target, pixel.V(16.0, 16.0), resources.manager)
	}
}

//LoadSpriteManager loads manager for sprites for resource
func (resources *PlayerResources) LoadSpriteManager(manager *sprite.Manager) {
	resources.manager = manager
}

//CreateCharacter creates character for game
func (resources *PlayerResources) CreateCharacter(name string, x int, y int) {
	if resources.manager != nil {
		var character basic.Character
		character.SetIndexForSprite(name, resources.manager)
		if character.IsValid() {
			character.SetPosition(x, y)
			resources.characters = append(resources.characters, character)
			resources.characterIndex = len(resources.characters) - 1
		}
	}
}

//GetActualCharacter gets actual game character
func (resources *PlayerResources) GetActualCharacter() *basic.Character {
	if resources.characterIndex < len(resources.characters) {
		return &resources.characters[resources.characterIndex]
	}
	return nil
}

//GetActualCharacterIndex gets actual character's index
func (resources *PlayerResources) GetActualCharacterIndex() int {
	return resources.characterIndex
}

//SetActualCharacterIndex sets actual character's index
func (resources *PlayerResources) SetActualCharacterIndex(index int) {
	if index > 0 && index < len(resources.characters) {
		resources.characterIndex = index
	} else {
		resources.characterIndex = 0
	}
}

//IncrementIndex increments index of actual character
func (resources *PlayerResources) IncrementIndex() {
	if resources.characterIndex == len(resources.characters)-1 {
		resources.characterIndex = 0
	} else {
		resources.characterIndex++
	}
}

//DecrementIndex decrements index of actual character
func (resources *PlayerResources) DecrementIndex() {
	if resources.characterIndex == 0 {
		resources.characterIndex = len(resources.characters) - 1
	} else {
		resources.characterIndex--
	}
}
