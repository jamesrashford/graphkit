package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jamesrashford/graphkit/io"
)

func main() {
	file, err := os.Open("examples/lollipop/graph_data.csv")
	if err != nil {
		panic(err)
	}

	var gio io.GraphIO = io.NewCSVIO("", "", "", "", true, true)

	graph, err := gio.ReadGraph(file)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	err = gio.WriteGraph(graph, buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
