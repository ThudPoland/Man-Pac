//Package game contains code for executing game
package game

import (
	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/level"
)

//Game is structure containing main data about actual game
type Game struct {
	levels     []level.Level
	levelIndex int
	resources  PlayerResources
	enemy      basic.Character
}
