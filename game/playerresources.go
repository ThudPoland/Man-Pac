package game

import (
	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//PlayerResources has info about all player resources
type PlayerResources struct {
	characters []basic.Character
	manager    *sprite.Manager
}

//Draw draws player resources
func (resources PlayerResources) Draw(target pixel.Target) {
	for element := range resources.characters {
		resources.characters[element].Draw(target, pixel.V(0.0, 0.0), resources.manager)
	}
}

//LoadSpriteManager loads manager for sprites for resource
func (resources *PlayerResources) LoadSpriteManager(manager *sprite.Manager) {
	resources.manager = manager
}
