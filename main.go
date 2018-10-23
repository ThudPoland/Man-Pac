package main

import (
	"github.com/ThudPoland/Man-Pac/level"
	"github.com/ThudPoland/Man-Pac/sprite"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	var manager sprite.Manager
	var level level.Level

	manager.LoadFromList("data/sprites/list.txt")
	level.Load("data/levels/level1.txt")

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
	level.Draw(win, manager)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
