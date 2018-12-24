package pathfinding

import "math"

type Graph struct {
	nodes map[struct {
		int
		int
	}]*Node
	width  int
	height int
}

func InitGraph() *Graph {
	data := new(Graph)
	return data
}

func (graph *Graph) FillGraph(layout [][]int) {
	graph.width = len(layout)
	graph.height = len(layout[graph.width])

	for x := range graph.width {
		for y := range graph.height {
			node := Node{distance: math.Inf(0),
				parent: nil, state: Unvisited}
			if layout[x][y] > 0 {
				node.state = Obstacle
			}
			graph.nodes[struct {
				int
				int
			}{x, y}] = node
		}
	}
}
