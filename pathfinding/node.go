package pathfinding

type Node struct {
	parent   *Node
	state    State
	distance int
	x        int
	y        int
}
