package pathfinding

//Graph is struct for pathfinding graph
type Graph struct {
	nodes  [][]*Node
	width  int
	height int
}

//InitGraph creates new graph
func InitGraph() *Graph {
	data := new(Graph)
	return data
}

//FillGraph is used to fill graph needed by pathfinding algorithms
func (graph *Graph) FillGraph(layout [][]int) {
	graph.height = len(layout)
	graph.width = len(layout[0])

	for y := range layout {
		for x := range layout[0] {
			node := &Node{distance: 0,
				parent: nil, state: Unvisited,
				x: x, y: y}
			if layout[x][y] > 0 {
				node.state = Obstacle
			}
			graph.nodes[x][y] = node
		}
	}
}
