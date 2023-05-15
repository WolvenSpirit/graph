package main

import (
	"fmt"

	gengraph "github.com/WolvenSpirit/graph/gen-graph"
)

const DefaultWeight = 1_000_000_000 // MAX WEIGHT

var v = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
}

var e = [][]int{
	// v1,v2,distance
	{0, 1, 30}, // a -> b
	{0, 2, -2}, // a -> c
	{2, 3, 21}, // c -> d
	{1, 4, 50}, // b -> e
	{3, 4, 11}, // d -> e
	{4, 5, 12}, // e -> f
	{4, 8, 30}, // d -> i
	{5, 8, 15}, // f -> i
}

type V = gengraph.V

type M map[string]V

type Options struct {
	StartVertex int
	Store       M
	Out         *[][]int
}

func initStoreWithVertexList(start int, v []string) []V {
	var vArr []V
	for k, val := range v {
		if k != start {
			vArr = append(vArr, V{Name: val, Index: k, Distance: DefaultWeight})
		} else {
			vArr = append(vArr, V{Name: val, Index: k, Distance: 0})
		}
	}
	return vArr
}

/*
		BellmanFord_FromVertexAndEdgeLists implementation details:

	  - Direction of the graph is given by the index value, higher index of v means v is downstream of lower index v
	  - Requires that the vertex list be parsed into an initial []V using initStoreWithVertexList
	    However this store is for internal use, the output of the call will be path which represents the shortest found path from start

		PS. Needs to be streamlined
*/
func BellmanFord_FromVertexAndEdgeLists(start int, v []string, e [][]int, store *[]V, path *[]V) []V {
	nextHop := V{Distance: DefaultWeight}
	s := (*store)
	for _, edge := range e {
		if edge[0] == start && edge[0] < edge[1] {
			s[edge[1]].Distance = s[edge[0]].Distance + edge[2]
			if s[edge[1]].Distance < nextHop.Distance {
				nextHop.Distance = s[edge[1]].Distance
				nextHop.Name = s[edge[1]].Name
				nextHop.Index = s[edge[1]].Index
			}
		}
		if edge[1] == start && edge[1] < edge[0] {
			s[edge[0]].Distance = s[edge[0]].Distance + edge[2]
			if s[edge[0]].Distance < nextHop.Distance {
				nextHop.Distance = s[edge[0]].Distance
				nextHop.Name = s[edge[0]].Name
				nextHop.Index = s[edge[0]].Index
			}
		}
	}
	if nextHop.Distance != DefaultWeight {
		(*path) = append((*path), nextHop)
		return BellmanFord_FromVertexAndEdgeLists(nextHop.Index, v, e, &s, path)
	}
	return *path
}

func main() {
	store := initStoreWithVertexList(0, v)
	var path []V
	path = append(path, V{Name: "a", Index: 0})
	out := BellmanFord_FromVertexAndEdgeLists(path[0].Index, v, e, &store, &path)
	totalDistance := 0
	for k, v := range out {
		var hopDistance int
		if k != 0 {
			hopDistance = (v.Distance - out[k-1].Distance)
			totalDistance += (v.Distance - out[k-1].Distance)
		}
		fmt.Printf("[Index: %d][Name: %s]<Distance: %d> ->\n", v.Index, v.Name, hopDistance)
	}
	fmt.Println("Total distance ", totalDistance)
	fmt.Println(gengraph.GenGraphToFile(out, "out.dot"))
}
