package io_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

// Read
func TestJSONIORead(t *testing.T) {
	jsonio := io.NewJSONIO()
	var readwrite io.GraphIO = jsonio

	buf := new(bytes.Buffer)
	buf.WriteString(`{"directed": false, "multigraph": false, "graph": {}, "nodes": [{"id": "0"}, {"id": "1"}, {"id": "2"}, {"id": "3"}], "links": [{"source": "0", "target": "1"},{"source": "0", "target": "2"},{"source": "0", "target": "3"},{"source": "1", "target": "2"},{"source": "1", "target": "3"},{"source": "2", "target": "3"}]}`)

	graph, err := readwrite.ReadGraph(buf)
	if err != nil {
		t.Error(err)
	}

	testGraph := models.NewEmptyGraph(false)
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

// Read Examples
func TestJSONIOExamples(t *testing.T) {
	paths, err := io.GetExamples("graph.json")
	if err != nil {
		t.Error(err)
	}

	jsonio := io.NewJSONIO()
	var readwrite io.GraphIO = jsonio

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

// Write
func TestJSONIOWrite(t *testing.T) {
	testGraph := models.NewEmptyGraph(true)
	testGraph.AddEdge("0", "1", nil)
	testGraph.AddEdge("0", "2", nil)
	testGraph.AddEdge("0", "3", nil)
	testGraph.AddEdge("1", "2", nil)
	testGraph.AddEdge("1", "3", nil)
	testGraph.AddEdge("2", "3", nil)

	jsonio := io.NewJSONIO()
	var readwrite io.GraphIO = jsonio

	buf := new(bytes.Buffer)

	err := readwrite.WriteGraph(testGraph, buf)
	if err != nil {
		t.Error(err)
	}

	resultGraph, err := readwrite.ReadGraph(buf)
	if err != nil {
		t.Error(err)
	}

	if !resultGraph.Equal(testGraph) {
		t.Errorf("test graph does not match read graph")
	}
}
