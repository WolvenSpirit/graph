package main

import "fmt"

type Vertex string

type Edge []int

var v = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", // g,h,j have no edges
}

var e = []Edge{
	//   v1,v2,distance
	[]int{0, 1, 30}, // a -> b
	[]int{0, 2, 20}, // a -> c
	[]int{2, 3, 21}, // c -> d
	[]int{1, 4, 50}, // b -> e
	[]int{3, 4, 11}, // d -> e
	[]int{4, 5, 12}, // e -> f
	[]int{4, 8, 30}, // e -> i
	[]int{5, 8, 15}, // f -> j
}

func main() {
	const addedDistanceOnConnect = 100
	var notLinkedVertexes []int
	// find index of vertex that is not linked
	for i := range v {
		found := false
		for _, edge := range e {
			if edge[0] == i || edge[1] == i {
				found = true
			}
		}
		if !found {
			notLinkedVertexes = append(notLinkedVertexes, i)
		}
	}
	// link the vertex to another vertex that has more than one edge
	vertexesEdgeCount := make(map[int]int, 0)
	edgesMinOfVertex := 2
	for _, edge := range e {
		addedConnection := false
		vertexesEdgeCount[edge[0]]++
		vertexesEdgeCount[edge[1]]++
		if vertexesEdgeCount[edge[0]] > edgesMinOfVertex && len(notLinkedVertexes) > 0 {
			edge[0] = notLinkedVertexes[0]
			edge[2] += addedDistanceOnConnect
			addedConnection = true
		} else if vertexesEdgeCount[edge[1]] > edgesMinOfVertex && len(notLinkedVertexes) > 0 {
			edge[1] = notLinkedVertexes[0]
			edge[2] += addedDistanceOnConnect
			addedConnection = true
		}
		if addedConnection {
			notLinkedVertexes = notLinkedVertexes[1:]
		}
	}
	fmt.Println(e)
}
