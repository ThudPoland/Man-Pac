package sprite

import (
	"bufio"
	"io"
	"os"
)

//Manager is manager for sprites in project
type Manager struct {
	sprites []Sprite
	counter int
}

//Add sprite to manager
func (manager *Manager) Add(path string) {
	var sprite Sprite
	sprite.Load(path, manager.counter)
	if sprite.spriteError == nil {
		sprite.id = manager.counter
		manager.counter++
		manager.sprites = append(manager.sprites, sprite)
	}
}

//Add sprites from .txt list
func (manager *Manager) LoadFromList(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		source, err := reader.ReadString('\n')
		if err == nil {
			manager.Add(source)
			continue
		} else if err == io.EOF {
			manager.Add(source)
			break
		} else {
			break
		}
	}
}
