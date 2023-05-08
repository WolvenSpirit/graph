/*
* Dijkstra algorithm
* Mihai Dragusin | dragusin.mihai.90@gmail.com
 */

package main

import (
	"fmt"
)

type Vertex string

type Edge []int

var v = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
}

var e = []Edge{
	//   v1,v2,distance
	[]int{0, 1, 30}, // a -> b
	[]int{0, 2, 20}, // a -> c
	[]int{2, 3, 21}, // c -> d
	[]int{1, 4, 50}, // b -> e
	[]int{3, 4, 11}, // d -> e
	[]int{4, 5, 12}, // d -> e
	[]int{4, 8, 30}, // d -> e
	[]int{5, 8, 15}, // d -> e
}

// Djkistra implementation
// - Where: i = first node index; previous will be set as -1 by caller, v = vertex array,
// e = edge array, out is the slice which stores the indexes, and edge distance to the shortest path
func Dijkstra(i int, previous int, v []string, e []Edge, out *[][]int) {
	var max int = 1e10
	var next int = -1
	var distance int
	for _, edge := range e {
		if edge[0] == i && max > edge[2] && edge[1] != previous {
			max = edge[2]
			next = edge[1]
			distance = edge[2]
		}
		if edge[1] == i && max > edge[2] && edge[0] != previous {
			max = edge[2]
			next = edge[0]
			distance = edge[2]
		}
	}
	if next == -1 || next < i {
		return
	}
	(*out) = append((*out), []int{i, next, distance})
	Dijkstra(next, i, v, e, out)
}

func main() {
	start := 0
	out := make([][]int, 0)
	Dijkstra(start, -1, v, e, &out)
	for k := range out {
		fmt.Printf("[%s] %s [%s] <%d>\n", v[out[k][0]], "->", v[out[k][1]], out[k][2])
	}
	// shortest distance
	var shortestPathDistance int
	for k := range out {
		shortestPathDistance += out[k][2]
	}
	fmt.Printf("---\nShortest distance to final linked node is %d\n", shortestPathDistance)
}
