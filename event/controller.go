package event

import (
	"github.com/faiface/pixel/pixelgl"
)

//Controller controls all events that are connected to game
type Controller struct {
	window  *pixelgl.Window
	actions map[pixelgl.Button]ActionData
}

//InitController inits controller
func (controller *Controller) InitController() {
	controller.actions = make(map[pixelgl.Button]ActionData)
}

//SetProcessedWindow sets a window to be processed by controller
func (controller *Controller) SetProcessedWindow(window *pixelgl.Window) {
	controller.window = window
}

//RegisterAction is used for registering actions.
func (controller *Controller) RegisterAction(button pixelgl.Button, receiver interface{}, action func(interface{})) {
	data := ActionData{receiver: receiver, action: action}
	controller.actions[button] = data
}

//ProcessEventLoop processes event loop
func (controller *Controller) ProcessEventLoop() {
	for key, data := range controller.actions {
		if controller.window.JustPressed(key) {
			data.Process()
		}
	}
}
