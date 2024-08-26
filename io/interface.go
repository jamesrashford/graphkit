package io

import "github.com/jamesrashford/graphkit/models"

type GraphIO interface {
	ReadGraph(path string, directed bool) (*models.Graph, error)
	WriteGraph(graph *models.Graph, path string) error
}
