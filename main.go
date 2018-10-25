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
	game.SetActualLevel(1)

	var manager sprite.Manager
	manager.LoadFromList("data/sprites/level/list.txt")

	game.SetLevelManager(&manager)

	cfg := pixelgl.WindowConfig{
		Title:  "Man-Pac The Game!",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Black)
	game.Draw(win)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
