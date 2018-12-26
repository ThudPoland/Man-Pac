package basic

//AIProcessable is processable interface for computer-driven characters
type AIProcessable interface {
	Processable
	DoCalculations()
}
