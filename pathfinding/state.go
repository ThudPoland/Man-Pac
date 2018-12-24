package pathfinding

package basic

//State defines state of node
type State int

const (
	//Unvisited means not visited node
	Unvisited State = iota
	//Visited means visited node
	Visited = iota
	//Explored means node visited and neighbour nodes also visiteda
	Explored = iota
	//Obstacle means node is obstacle for another nodes
	Obstacle = iota
)