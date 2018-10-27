package event

import (
	"github.com/faiface/pixel/pixelgl"
)

//Controller controls all events that are connected to game
type Controller struct {
	window  *pixelgl.Window
	actions map[pixelgl.Button]ActionData
}

//SetProcessedWindow sets a window to be processed by controller
func (controller *Controller) SetProcessedWindow(window *pixelgl.Window) {
	controller.window = window
}

func (controller *Controller) RegisterAction()
