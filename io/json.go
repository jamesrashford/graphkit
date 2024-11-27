package io

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jamesrashford/graphkit/models"
)

type JSONIO struct {
}

func NewJSONIO() *JSONIO {
	return &JSONIO{}
}

func (jsonio *JSONIO) ReadGraph(reader io.Reader) (*models.Graph, error) {
	/*
		b, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}*/

	decoder := json.NewDecoder(reader)

	var jsonGraph models.JSONGraph
	err := decoder.Decode(&jsonGraph) // force int to string
	if err != nil {
		return nil, err
	}

	graph := models.NewEmptyGraph(jsonGraph.Directed)

	for _, n := range jsonGraph.Nodes {
		// TODO: Include attrs
		id := fmt.Sprintf("%v", n.ID)
		graph.AddNode(id)
	}

	for _, e := range jsonGraph.Links {
		// TODO: Include attrs
		s := fmt.Sprintf("%v", e.Source)
		t := fmt.Sprintf("%v", e.Target)
		graph.AddEdge(s, t, nil)
	}

	return graph, nil
}

func (jsonio *JSONIO) WriteGraph(graph *models.Graph, writer io.Writer) error {
	// Load our graph into JSONGraph struct
	jg := models.JSONGraph{}
	jg.Directed = graph.Directed
	jg.Multigraph = false

	for _, n := range graph.GetNodes() {
		node := struct {
			ID interface{} `json:"id"`
		}{
			ID: n.ID,
		}
		jg.Nodes = append(jg.Nodes, node)
	}

	for _, e := range graph.GetEdges() {
		edge := struct {
			Source interface{} `json:"source"`
			Target interface{} `json:"target"`
		}{
			Source: e.Source.ID,
			Target: e.Target.ID,
		}
		jg.Links = append(jg.Links, edge)
	}

	// Write JSONGraph to buffer
	data, err := json.MarshalIndent(jg, "", " ")
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
