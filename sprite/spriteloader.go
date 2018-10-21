package sprite

//Loader allows sprite loading for game
type Loader interface {
	Load(path string)
}

/*//Load loads image to sprite
func (sprite Sprite) Load(path string) {
	sprite.picture = nil
	file, err := os.Open(path)
}*/
