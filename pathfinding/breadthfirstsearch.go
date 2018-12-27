package pathfinding

import (
	"github.com/ThudPoland/Man-Pac/basic"
)

//BreadthFirstSearch is a struct for making pathfinding
type BreadthFirstSearch struct {
	lastNode  *Node
	direction basic.Direction
	reachable bool
}

//GetSearchResult creates result for Breadth First Search
func (searchData *BreadthFirstSearch) GetSearchResult(startX, startY, endX, endY int, source *Graph) {
	searchData.reachable = false
	node := source.nodes[startX][startY]
	nodesSlice := []*Node{node}
	for len(nodesSlice) > 0 {
		actualNode := nodesSlice[0]
		nodesSlice = nodesSlice[1:]
		neighbours := searchData.getNeighbours(startX, startY, source)

		for element := range neighbours {
			actualNeighbour := neighbours[element]
			if actualNeighbour.state == Unvisited {
				actualNeighbour.state = Visited
				actualNeighbour.parent = actualNode
				actualNeighbour.distance = actualNeighbour.parent.distance + 1

				if actualNeighbour.x == endX && actualNeighbour.y == endY {
					searchData.reachable = true
					searchData.lastNode = actualNeighbour
					return
				}

				nodesSlice = append(nodesSlice, actualNeighbour)
			}
		}

		actualNode.state = Explored
	}
}

func (searchData *BreadthFirstSearch) getNeighbours(startX, startY int, source *Graph) []*Node {
	neighbours := []*Node{}

	coordinates := []struct{ x, y int }{{startX - 1, startY}, {startX + 1, startY}, {startX, startY - 1}, {startX, startY + 1}}
	for element := range coordinates {
		if coordinates[element].x < 0 || coordinates[element].x >= source.width {
			if coordinates[element].y < 0 || coordinates[element].y >= source.height {
				continue
			}
		}

		node := source.nodes[coordinates[element].x][coordinates[element].y]
		if node.state == Unvisited {
			neighbours = append(neighbours, node)
		}
	}

	return neighbours
}

func (searchData *BreadthFirstSearch) getFirstNode(source *Graph, sourceNode *Node) {
	if searchData.lastNode != nil && searchData.reachable == true {
		for {
			actualNode := sourceNode
			if actualNode == searchData.lastNode {
				searchData.direction = basic.No
			}
		}
	}
}
