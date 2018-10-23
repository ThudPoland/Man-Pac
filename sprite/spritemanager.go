package sprite

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/faiface/pixel"
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

//LoadFromList function adds sprites from .txt list
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
			manager.Add(strings.Trim(source, "\n"))
			continue
		} else if err == io.EOF {
			manager.Add(strings.Trim(source, "\n"))
			break
		} else {
			break
		}
	}
}

//GetSprite is function for getting sprite for usage
func (manager Manager) GetSprite(index int) *Sprite {
	if index < len(manager.sprites) {
		return &manager.sprites[index]
	}
	return nil
}

//DrawSprite draws sprite in batch
func (manager Manager) DrawSprite(index int, target pixel.Target, destination pixel.Matrix) {
	if index < len(manager.sprites) {
		manager.sprites[index].Draw(target, destination)
	}
}
