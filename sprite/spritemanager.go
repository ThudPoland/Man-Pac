package sprite

//Manager is manager for sprites in project
type Manager struct {
	sprites []Sprite
	counter int
}

//Add sprite to manager
func (manager Manager) Add(path string) {
	var sprite Sprite
	sprite.Load(path, manager.counter)
	manager.counter++
	manager.sprites = append(manager.sprites, sprite)
}
