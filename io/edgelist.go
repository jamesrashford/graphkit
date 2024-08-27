package io

import (
	"bufio"
	"io"
	"strings"

	"github.com/jamesrashford/graphkit/models"
)

type EdgeListIO struct {
	Comments  string
	Delimiter string
	Directed  bool
}

func NewEdgeListIO(comment, delimiter string, directed bool) *EdgeListIO {
	if comment == "" {
		comment = "#"
	}
	if delimiter == "" {
		delimiter = " "
	}

	elio := EdgeListIO{
		Comments:  comment,
		Delimiter: delimiter,
		Directed:  directed,
	}
	return &elio
}

// data, if true, encode as json and store in params

func (elio *EdgeListIO) ReadGraph(reader *io.Reader) (*models.Graph, error) {
	graph := models.NewEmptyGraph(elio.Directed)

	/*
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	*/

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, elio.Comments) {
			continue
		}

		row := strings.Split(line, elio.Delimiter)

		// check for json data in idx 2. If so, load into params
		source := row[0]
		target := row[1]

		//check if more
		//data := row[2]

		graph.AddEdge(source, target, nil)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func (elio *EdgeListIO) WriteGraph(graph *models.Graph, writer io.Writer) {
	//Convert String id to numerial id
	//nodeIDMap := make(map[string]int)
	return
}
