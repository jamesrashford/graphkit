package main

import (
	"fmt"

	"github.com/jamesrashford/GoGraph/models"
)

func main() {
	graph := models.NewEmptyGraph(true)
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("E", "F")

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
}
