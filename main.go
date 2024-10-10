package main

import (
	"math"
	"os"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/layout"
	"github.com/jamesrashford/graphkit/plot"
)

func main() {
	elio := io.NewEdgeListIO("", "", true)

	file, err := os.Open("examples/complete/graph.edgelist")
	if err != nil {
		panic(err)
	}

	G, err := elio.ReadGraph(file)
	if err != nil {
		panic(err)
	}

	k := math.Sqrt(1.0 / float64(G.NoNodes))
	pos := layout.ForceDirected(G, nil, 50, k, 0.1)

	plt := plot.NewGraphPlotter(800, 600)
	plt.Draw(G, pos, true, "test_plot.png")
}
