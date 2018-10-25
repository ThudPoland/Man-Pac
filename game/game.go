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
		game.resources.Draw(t)
	}
}

//SetLevelManager sets level sprites manager
func (game *Game) SetLevelManager(spriteManager *sprite.Manager) {
	game.levelManager = spriteManager
}

//SetResourcesManager sets character sprites manager
func (game *Game) SetResourcesManager(spriteManager *sprite.Manager) {
	game.resources.LoadSpriteManager(spriteManager)
}

//LoadLevel implements level loading for game
func (game *Game) LoadLevel(path string) {
	var level level.Level
	level.Load(path)
	level.SetOffset(pixel.V(0.0, 0.0))
	game.levels = append(game.levels, level)
}

//AddGhostToLevel adds ghost to level
func (game *Game) AddGhostToLevel(x int, y int) {
	game.resources.CreateCharacter("ghost", x, y)
}
