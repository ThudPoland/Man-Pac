package sprite

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

//Loader allows sprite loading for game
type Loader interface {
	Load(path string, id int)
}

//Load loads image to sprite
func (sprite Sprite) Load(path string, id int) {
	sprite.picture = nil
	sprite.id = id
	file, err := os.Open(path)

	if err != nil {
		sprite.picture = nil
		sprite.spriteError = err
		return
	}
	defer file.Close()

	image, _, err := image.Decode(file)

	if err != nil {
		sprite.picture = nil
		sprite.spriteError = err
		return
	}

	sprite.picture = pixel.PictureDataFromImage(image)
	sprite.sprite = pixel.NewSprite(sprite.picture, sprite.picture.Bounds())
}
