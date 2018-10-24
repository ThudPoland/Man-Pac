package game

//Loader provides interface for loading game parts
type Loader interface {
	LoadLevel(path string)
}
