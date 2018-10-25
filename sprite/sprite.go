//Package sprite provides sprites for game
package sprite

import (
	"image/png"
	"os"

	"github.com/faiface/pixel"
)

//Sprite implements pixel's image as a part of game
type Sprite struct {
	picture     pixel.Picture
	sprite      *pixel.Sprite
	spriteError error
	id          int
	name        string
}

//Draw function draws sprite
func (sprite Sprite) Draw(target pixel.Target, destination pixel.Matrix) {
	if sprite.sprite != nil {
		sprite.sprite.Draw(target, destination)
	}
}

//Load loads image to sprite
func (sprite *Sprite) Load(path string, id int) {
	sprite.picture = nil
	sprite.id = id
	sprite.spriteError = nil
	file, err := os.Open(path)

	if err != nil {
		sprite.picture = nil
		sprite.spriteError = err
		return
	}
	defer file.Close()

	image, err := png.Decode(file)

	if err != nil {
		sprite.picture = nil
		sprite.spriteError = err
		return
	}

	sprite.picture = pixel.PictureDataFromImage(image)
	sprite.sprite = pixel.NewSprite(sprite.picture, sprite.picture.Bounds())
}
