package main

import (
	"bytes"
	"fmt"

	"github.com/jamesrashford/graphkit/io"
)

func main() {
	elio := io.NewEdgeListIO("", "", true)

	buf := new(bytes.Buffer)
	buf.WriteString("0 1 5\n0 2 6\n0 3 2\n0 4 1\n1 2 8\n1 3 3\n1 4 4\n2 3 9\n2 4 5\n3 4 1")

	fmt.Println(buf.String())

	graph, err := elio.ReadGraph(buf)
	if err != nil {
		panic(err)
	}

	for _, e := range graph.GetEdges() {
		fmt.Println(e)
	}

	buf = new(bytes.Buffer)

	elio.WriteGraph(graph, buf)

	out := buf.String()
	fmt.Println(out)

}
