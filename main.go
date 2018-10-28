package main

import (
	"github.com/ThudPoland/Man-Pac/game"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	var game game.Game
	game.LoadLevel("data/levels/level1.txt")

	var levelManager sprite.Manager
	levelManager.LoadFromList("data/sprites/level/list.txt")
	var actorsManager sprite.Manager
	actorsManager.LoadFromList("data/sprites/characters/list.txt")
	var interfaceManager sprite.Manager
	interfaceManager.LoadFromList("data/sprites/ui/list.txt")

	game.SetLevelManager(&levelManager)
	game.SetResourcesManager(&actorsManager)
	game.SetInterfaceManager(&interfaceManager)
	game.AddGhostToLevel(3, 11)
	game.AddGhostToLevel(4, 11)
	game.AddGhostToLevel(5, 11)
	game.SetActualLevel(1)

	cfg := pixelgl.WindowConfig{
		Title:  "Man-Pac The Game!",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Black)
		game.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
