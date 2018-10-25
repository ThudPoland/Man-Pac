package sprite

//Loader allows sprite loading for game
type Loader interface {
	Load(path string, id int)
}
