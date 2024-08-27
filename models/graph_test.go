package models_test

import (
	"testing"

	"github.com/jamesrashford/graphkit/models"
)

func TestGraphEquality(t *testing.T) {

	g1 := models.NewEmptyGraph(true)
	g2 := models.NewEmptyGraph(true)

	want := true
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are meant to be equal. Expected %v, got %v", want, res)
	}

	g1 = models.NewEmptyGraph(true)
	g2 = models.NewEmptyGraph(false)

	want = false
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are not meant to be equal. Expected %v, got %v", want, res)
	}

	g1 = models.NewEmptyGraph(true)
	g1.AddEdge("A", "B", nil)
	g1.AddEdge("B", "C", nil)
	g1.AddEdge("C", "D", nil)

	g2 = models.NewEmptyGraph(true)
	g2.AddEdge("A", "B", nil)
	g2.AddEdge("B", "C", nil)
	g2.AddEdge("C", "D", nil)

	want = true
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are meant to be equal. Expected %v, got %v", want, res)
	}

	g1 = models.NewEmptyGraph(true)
	g1.AddEdge("A", "B", nil)
	g1.AddEdge("B", "C", nil)
	g1.AddEdge("C", "D", nil)
	g1.AddEdge("D", "E", nil)

	g2 = models.NewEmptyGraph(true)
	g2.AddEdge("A", "B", nil)
	g2.AddEdge("B", "C", nil)
	g2.AddEdge("C", "D", nil)

	want = false
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are not meant to be equal. Expected %v, got %v", want, res)
	}

	g1 = models.NewEmptyGraph(true)
	g1.AddEdge("A", "B", nil)
	g1.AddEdge("B", "C", nil)
	g1.AddEdge("C", "D", nil)

	g2 = models.NewEmptyGraph(true)
	g2.AddEdge("A", "B", nil)
	g2.AddEdge("B", "C", nil)
	g2.AddEdge("C", "F", nil)

	want = false
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are not meant to be equal. Expected %v, got %v", want, res)
	}

	g1 = models.NewEmptyGraph(true)
	g1.AddEdge("1", "2", nil)
	g1.AddEdge("2", "3", nil)
	g1.AddEdge("3", "4", nil)

	g2 = models.NewEmptyGraph(true)
	g2.AddEdge("A", "B", nil)
	g2.AddEdge("B", "C", nil)
	g2.AddEdge("C", "D", nil)

	want = false
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are not meant to be equal. Expected %v, got %v", want, res)
	}

	g1 = models.NewEmptyGraph(true)
	g1.AddNode("A")
	g1.AddNode("B")
	g1.AddNode("C")
	g1.AddNode("D")

	g2 = models.NewEmptyGraph(true)
	g2.AddNode("1")
	g2.AddNode("2")
	g2.AddNode("3")
	g2.AddNode("4")

	want = false
	if res := g1.Equal(g2); res != want {
		t.Errorf("Graphs are not meant to be equal. Expected %v, got %v", want, res)
	}
}
