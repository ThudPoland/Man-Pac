//Package game contains code for executing game
package game

import (
	"github.com/ThudPoland/Man-Pac/basic"
	"github.com/ThudPoland/Man-Pac/level"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
)

//Game is structure containing main data about actual game
type Game struct {
	levels         []level.Level
	levelIndex     int
	resources      PlayerResources
	enemy          basic.Character
	levelManager   *sprite.Manager
	gameSpriteSize int
}

//SetActualLevel sets index of a level
func (game *Game) SetActualLevel(index int) {
	game.levelIndex = index
}

//Draw draws entire actual game
func (game Game) Draw(t pixel.Target) {
	index := game.levelIndex - 1
	if index >= 0 && index < len(game.levels) {
		game.levels[index].Draw(t, *game.levelManager)
		for element := range game.resources.characters {
			game.resources.characters[element].Draw(t, pixel.V(0.0, 0.0), nil)
		}
	}
}

//SetLevelManager sets level sprites manager
func (game *Game) SetLevelManager(spriteManager *sprite.Manager) {
	game.levelManager = spriteManager
}

//LoadLevel implements level loading for game
func (game *Game) LoadLevel(path string) {
	var level level.Level
	level.Load(path)
	level.SetOffset(pixel.V(0.0, 0.0))
	game.levels = append(game.levels, level)
}
