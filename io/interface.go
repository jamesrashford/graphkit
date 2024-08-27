package io

import (
	"io"

	"github.com/jamesrashford/graphkit/models"
)

type GraphIO interface {
	ReadGraph(reader io.Reader) (*models.Graph, error)
	WriteGraph(graph *models.Graph, writer io.Writer) error
}
