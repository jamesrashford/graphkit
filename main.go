package main

import (
	"fmt"

	"github.com/jamesrashford/graphkit/models"
)

func main() {
	graph := models.NewEmptyGraph(true)
	graph.AddEdge("A", "B", nil)
	graph.AddEdge("B", "C", nil)
	graph.AddEdge("C", "D", nil)
	graph.AddEdge("C", "E", nil)
	graph.AddEdge("E", "F", nil)

	fmt.Println("Nodes")
	nodes := graph.GetNodes()
	for _, node := range nodes {
		fmt.Println(node.ID, node.Params)
	}

	fmt.Println("Edges")
	edges := graph.GetEdges()
	for _, edge := range edges {
		fmt.Println(edge.Source.ID, edge.Target.ID)
	}

	//path := "examples/complete/graph.edgelist"
	//var gio io.GraphIO
}
