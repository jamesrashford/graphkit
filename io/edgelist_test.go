package io_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

// TEST Read and write expected results with bytes buffer
func TestEdglistIORead(t *testing.T) {
	elio := io.NewEdgeListIO("", "", true)

	buf := new(bytes.Buffer)
	buf.WriteString("0 1\n0 2\n0 3\n1 2\n1 3\n2 3")

	graph, err := elio.ReadGraph(buf)
	if err != nil {
		t.Error(err)
	}

	testGraph := models.NewEmptyGraph(true)
	testGraph.AddEdge("0", "1", nil)
	testGraph.AddEdge("0", "2", nil)
	testGraph.AddEdge("0", "3", nil)
	testGraph.AddEdge("1", "2", nil)
	testGraph.AddEdge("1", "3", nil)
	testGraph.AddEdge("2", "3", nil)

	if !graph.Equal(testGraph) {
		t.Errorf("test graph does not match read graph")
	}

}

func TestEdglistIOWrite(t *testing.T) {
	testGraph := models.NewEmptyGraph(true)
	testGraph.AddEdge("0", "1", nil)
	testGraph.AddEdge("0", "2", nil)
	testGraph.AddEdge("0", "3", nil)
	testGraph.AddEdge("1", "2", nil)
	testGraph.AddEdge("1", "3", nil)
	testGraph.AddEdge("2", "3", nil)

	elio := io.NewEdgeListIO("", "", true)

	buf := new(bytes.Buffer)

	fmt.Println(buf.String())

	err := elio.WriteGraph(testGraph, buf)
	if err != nil {
		t.Error(err)
	}

	expected := "0 1\n0 2\n0 3\n1 2\n1 3\n2 3\n"
	ans := buf.String()

	if strings.Compare(expected, ans) != 0 {
		t.Errorf("test graph does not match write graph")
	}
}

// TEST with one col and two

// TEST with and without weight
