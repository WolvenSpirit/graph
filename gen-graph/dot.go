package gengraph

import (
	"fmt"
	"os"
	"strconv"

	"github.com/awalterschulze/gographviz"
)

type V struct {
	Name     string
	Index    int
	Distance int
}

/*
*	GenGraphToFile writes a graph to file from the path output []V
 */
func GenGraphToFile(v []V, fileName string) error {
	gAst, _ := gographviz.ParseString(`digraph G {}`)
	g := gographviz.NewGraph()
	gographviz.Analyse(gAst, g)
	for k := range v {
		err := g.AddNode("G", v[k].Name, nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		if k != 0 {
			a := make(map[string]string, 0)
			a["label"] = strconv.Itoa(v[k].Distance - v[k-1].Distance)
			err := g.AddEdge(v[k-1].Name, v[k].Name, true, a)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	return os.WriteFile(fileName, []byte(g.String()), os.ModePerm)
}
