package io_test

import (
	"bytes"
	"testing"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

// TODO
func TestGraphologyIORead(t *testing.T) {
	gio := io.NewGraphologyIO()
	var readwrite io.GraphIO = gio

	buf := new(bytes.Buffer)
	buf.WriteString(`{"attributes":{"name":"My Graph"},"options":{"type":"directed"},"nodes":[{"key":"0"},{"key":"1"},{"key":"2"},{"key":"3"}],"edges":[{"key":"0->1","source":"0","target":"1"},{"key":"0->2","source":"0","target":"2"},{"key":"0->3","source":"0","target":"3"},{"key":"1->2","source":"1","target":"2"},{"key":"1->3","source":"1","target":"3"},{"key":"2->3","source":"2","target":"3"}]}`)

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

func TestGraphologyIOWrite(t *testing.T) {
	testGraph := models.NewEmptyGraph(true)
	testGraph.AddEdge("0", "1", nil)
	testGraph.AddEdge("0", "2", nil)
	testGraph.AddEdge("0", "3", nil)
	testGraph.AddEdge("1", "2", nil)
	testGraph.AddEdge("1", "3", nil)
	testGraph.AddEdge("2", "3", nil)

	gio := io.NewGraphologyIO()
	var readwrite io.GraphIO = gio

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

// TODO: TEST with attributes
