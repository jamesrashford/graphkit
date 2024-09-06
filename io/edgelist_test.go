package io_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

func TestEdglistIORead(t *testing.T) {
	elio := io.NewEdgeListIO("", "", true)
	var readwrite io.GraphIO = elio

	buf := new(bytes.Buffer)
	buf.WriteString("0 1\n0 2\n0 3\n1 2\n1 3\n2 3")

	graph, err := readwrite.ReadGraph(buf)
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

func TestEdglistIOExamples(t *testing.T) {
	paths, err := io.GetExamples(".edgelist")
	if err != nil {
		t.Error(err)
	}

	elio := io.NewEdgeListIO("", "", true)
	var readwrite io.GraphIO = elio

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			file, err := os.Open(path)
			if err != nil {
				t.Error(err)
			}

			defer file.Close()
			_, err = readwrite.ReadGraph(file)
			if err != nil {
				t.Error(err)
			}
		})
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
	var readwrite io.GraphIO = elio

	buf := new(bytes.Buffer)

	fmt.Println(buf.String())

	err := readwrite.WriteGraph(testGraph, buf)
	if err != nil {
		t.Error(err)
	}

	expected := "0 1\n0 2\n0 3\n1 2\n1 3\n2 3\n"
	ans := buf.String()

	if strings.Compare(expected, ans) != 0 {
		t.Errorf("test graph does not match write graph")
	}
}

// TEST with and without weight
func TestEdglistIOWeight(t *testing.T) {
	elio := io.NewEdgeListIO("", "", true)
	var readwrite io.GraphIO = elio

	input := "0 1 5\n0 2 3\n0 3 7\n1 2 5\n1 3 7\n2 3 1\n"

	buf1 := new(bytes.Buffer)
	buf1.WriteString(input)

	graph, err := readwrite.ReadGraph(buf1)
	if err != nil {
		t.Error(err)
	}

	buf2 := new(bytes.Buffer)
	err = readwrite.WriteGraph(graph, buf2)
	if err != nil {
		t.Error(err)
	}

	output := buf2.String()

	if strings.Compare(input, output) != 0 {
		t.Errorf("weighted graphs do not match")
	}
}

func TestEdglistIOCols(t *testing.T) {
	elio := io.NewEdgeListIO("", "", true)
	var readwrite io.GraphIO = elio

	input := "0\n1\n2\n"

	buf := new(bytes.Buffer)
	buf.WriteString(input)

	_, err := readwrite.ReadGraph(buf)
	if err == nil {
		t.Errorf("expected error as not enough cols")
	}

	input = "0 2 3 4 5\n\n2 3 4 5 6\n"

	buf = new(bytes.Buffer)
	buf.WriteString(input)

	_, err = readwrite.ReadGraph(buf)
	if err == nil {
		t.Errorf("expected error as too many cols")
	}
}
