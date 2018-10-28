package main

import (
	"github.com/ThudPoland/Man-Pac/event"
	"github.com/ThudPoland/Man-Pac/game"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	var gameProvider game.Game
	gameProvider.LoadLevel("data/levels/level1.txt")
	var controller event.Controller

	var levelManager sprite.Manager
	levelManager.LoadFromList("data/sprites/level/list.txt")
	var actorsManager sprite.Manager
	actorsManager.LoadFromList("data/sprites/characters/list.txt")
	var interfaceManager sprite.Manager
	interfaceManager.LoadFromList("data/sprites/ui/list.txt")

	gameProvider.SetLevelManager(&levelManager)
	gameProvider.SetResourcesManager(&actorsManager)
	gameProvider.SetInterfaceManager(&interfaceManager)
	gameProvider.AddGhostToLevel(3, 11)
	gameProvider.AddGhostToLevel(4, 11)
	gameProvider.AddGhostToLevel(5, 11)
	gameProvider.SetActualLevel(1)

	cfg := pixelgl.WindowConfig{
		Title:  "Man-Pac The Game!",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)

	controller.InitController()
	controller.SetProcessedWindow(win)
	controller.RegisterAction(pixelgl.KeyQ, &gameProvider, game.PreviousCharacter)
	controller.RegisterAction(pixelgl.KeyE, &gameProvider, game.NextCharacter)

	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Black)
		controller.ProcessEventLoop()
		gameProvider.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
