package basic

//Processable implements processing functions for game
type Processable interface {
	IsReady() bool
	ProcessTurn()
}
