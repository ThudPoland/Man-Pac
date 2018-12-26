//Package game contains code for executing game
package game

import (
	"github.com/ThudPoland/Man-Pac/ai"
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
	enemy            *basic.ManPac
	levelManager     *sprite.Manager
	interfaceManager *sprite.Manager
	spriteManager    *sprite.Manager
	gameSpriteSize   int
	foodArray        level.FoodArray
}

//SetActualLevel sets index of a level
func (game *Game) SetActualLevel(index int) {
	game.levelIndex = index
	game.foodArray.GenerateFoodArray(&game.levels[game.levelIndex-1], game.resources.characters)
	game.SetEnemyInLevel(7, 11)
}

//Draw draws entire actual game
func (game *Game) Draw(t pixel.Target) {
	index := game.levelIndex - 1
	if index >= 0 && index < len(game.levels) {
		game.levels[index].Draw(t, *game.levelManager)
		game.foodArray.Draw(t, *game.levelManager)
		game.resources.Draw(t)
		if game.enemy != nil {
			game.enemy.Draw(t, pixel.V(16.0, 16.0), game.spriteManager)
		}
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
	game.spriteManager = spriteManager
	game.resources.LoadSpriteManager(spriteManager)
}

//SetInterfaceManager sets interface sprites manager
func (game *Game) SetInterfaceManager(spriteManager *sprite.Manager) {
	game.interfaceManager = spriteManager
	game.resources.LoadInterfaceSpriteManager(spriteManager)
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
	game.resources.CreateGhost("ghost", x, y)
}

//SetEnemyInLevel sets ManPac
func (game *Game) SetEnemyInLevel(x int, y int) {
	game.enemy = &basic.ManPac{}
	algorithm := &ai.Dumb{}
	algorithm.SetInput(&game.levels[game.levelIndex-1])
	game.enemy.SetAI(algorithm)
	game.enemy.SetPosition(x, y)
	game.enemy.SetSpriteManager(game.spriteManager)
	game.enemy.SetIndexForSprite("pacman", game.spriteManager)
}

//SetDirection sets direction for player
func (game *Game) SetDirection(direction basic.Direction) {
	character, ok := game.resources.GetActualCharacter().(*basic.Ghost)
	if ok == true {
		index := game.levelIndex - 1
		points := game.levels[index].GetNearPoints(character.GetX(), character.GetY())
		character.SetDirection(direction, points)
	}
}

//ProcessTurn is used for entire game to process turn
func (game *Game) ProcessTurn() {
	if game.resources.AreResourcesReady() {
		for element := range game.resources.characters {
			game.resources.characters[element].ProcessTurn()
		}
		game.enemy.DoCalculations()
		game.enemy.ProcessTurn()
		game.foodArray.Eat(game.enemy.GetX(), game.enemy.GetY())
	}
}

/*func (game *Game) setEnemyDirection() {
	index := game.levelIndex - 1
	possibleDirections := []basic.Direction{basic.Left, basic.Down, basic.Up, basic.Right}
	points := game.levels[index].GetNearPoints(game.enemy.GetX(), game.enemy.GetY())
	game.enemy.SetDirection(game.enemy.GetDirection(), points)
	if game.enemy.GetDirection() == basic.No || points.Sum() < 2 {
		for {
			value := rand.Intn(4)
			game.enemy.SetDirection(possibleDirections[value], points)
			if game.enemy.GetDirection() != basic.No {
				break
			}
		}
	}
}*/
