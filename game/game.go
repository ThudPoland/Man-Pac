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
	levels           []level.Level
	levelIndex       int
	resources        PlayerResources
	enemy            basic.Character
	levelManager     *sprite.Manager
	interfaceManager *sprite.Manager
	gameSpriteSize   int
	foodArray        level.FoodArray
}

//SetActualLevel sets index of a level
func (game *Game) SetActualLevel(index int) {
	game.levelIndex = index
	game.foodArray.GenerateFoodArray(&game.levels[game.levelIndex-1], game.resources.characters)
}

//Draw draws entire actual game
func (game Game) Draw(t pixel.Target) {
	index := game.levelIndex - 1
	if index >= 0 && index < len(game.levels) {
		game.levels[index].Draw(t, *game.levelManager)
		game.resources.Draw(t)
		game.foodArray.Draw(t, *game.levelManager)
		actualCharacter := game.resources.GetActualCharacter()
		if actualCharacter != nil {
			spriteSize := game.interfaceManager.GetSpriteSize()
			selectionPosition := pixel.IM.Moved(pixel.V(16.0, 16.0)).
				Moved(pixel.V(float64(actualCharacter.GetX()*spriteSize),
					float64(actualCharacter.GetY()*spriteSize)))
			game.interfaceManager.DrawSprite(game.interfaceManager.GetSpriteIndexByName("selection"),
				t, selectionPosition)
		}
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

func (game *Game) SetInterfaceManager(spriteManager *sprite.Manager) {
	game.interfaceManager = spriteManager
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
