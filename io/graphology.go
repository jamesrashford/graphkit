package io

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jamesrashford/graphkit/models"
)

type GraphologyIO struct {
}

func NewGraphologyIO() *GraphologyIO {
	return &GraphologyIO{}
}

func (graphologyio *GraphologyIO) ReadGraph(reader io.Reader) (*models.Graph, error) {
	decoder := json.NewDecoder(reader)

	var gg models.GraphologyGraph
	err := decoder.Decode(&gg)
	if err != nil {
		return nil, err
	}

	directed := false
	if val, ok := gg.Options["type"]; ok {
		if val == "directed" {
			directed = true
		}
	}

	graph := models.NewEmptyGraph(directed)
	graph.Params = gg.Attributes

	for _, n := range gg.Nodes {
		graph.AddNode(n.Key)
	}

	for _, e := range gg.Edges {
		graph.AddEdge(e.Source, e.Target, e.Attributes)
	}

	return graph, nil
}

func (graphologyio *GraphologyIO) WriteGraph(graph *models.Graph, writer io.Writer) error {
	gg := models.GraphologyGraph{}
	gg.Options = make(map[string]interface{})
	if graph.Directed {
		gg.Options["type"] = "directed"
	}
	gg.Attributes = graph.Params

	for _, n := range graph.GetNodes() {
		node := struct {
			Key        string                 `json:"key"`
			Attributes map[string]interface{} `json:"attributes"`
		}{
			Key:        n.ID,
			Attributes: n.Params,
		}
		gg.Nodes = append(gg.Nodes, node)
	}

	for _, e := range graph.GetEdges() {
		edge := struct {
			Key        string                 `json:"key"`
			Source     string                 `json:"source"`
			Target     string                 `json:"target"`
			Attributes map[string]interface{} `json:"attributes"`
		}{
			Key:        fmt.Sprintf("%s->%s", e.Source.ID, e.Target.ID),
			Source:     e.Source.ID,
			Target:     e.Target.ID,
			Attributes: e.Params,
		}
		gg.Edges = append(gg.Edges, edge)
	}

	data, err := json.MarshalIndent(gg, "", " ")
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
