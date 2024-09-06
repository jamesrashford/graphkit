package io_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

func TestCSVIORead(t *testing.T) {
	csvio := io.NewCSVIO("#", ",", "source", "target", true)
	var readwrite io.GraphIO = csvio

	buf := new(bytes.Buffer)
	buf.WriteString("source,target\n0,1\n0,2\n0,3\n1,2\n1,3\n2,3")

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

// Test Read Examples
func TestCSVIOExamples(t *testing.T) {
	paths, err := io.GetExamples(".csv")
	if err != nil {
		t.Error(err)
	}

	csvio := io.NewCSVIO("#", ",", "source", "target", true)
	var readwrite io.GraphIO = csvio

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

// Test Write

// Test params

// Test cols

// Test col names
