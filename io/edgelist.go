package io

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

func (elio *EdgeListIO) ReadGraph(reader io.Reader) (*models.Graph, error) {
	graph := models.NewEmptyGraph(elio.Directed)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, elio.Comments) {
			continue
		}

		row := strings.Split(line, elio.Delimiter)
		no_rows := len(row)
		if no_rows < 2 || no_rows > 3 {
			return nil, fmt.Errorf("need 2 or 3 cols, not %d", len(row))
		}

		// check for json data in idx 2. If so, load into params
		source := row[0]
		target := row[1]

		var params map[string]interface{}

		if len(row) > 2 {
			// Treat third col as weight
			weight := row[2]
			if val, err := strconv.Atoi(weight); err == nil {
				params = make(map[string]interface{})
				params["weight"] = val
			}
		}

		graph.AddEdge(source, target, params)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func (elio *EdgeListIO) WriteGraph(graph *models.Graph, writer io.Writer) error {
	//Convert String id to numerial id
	nodeIDMap := make(map[string]int)

	// Enumerate in order
	for i, node := range graph.GetNodes() {
		nodeIDMap[node.ID] = i
	}

	w := bufio.NewWriter(writer)

	for _, edge := range graph.GetEdges() {
		source := nodeIDMap[edge.Source.ID]
		target := nodeIDMap[edge.Target.ID]

		line := fmt.Sprintf("%d %d", source, target)

		if weight, ok := edge.Params["weight"]; ok {
			line += fmt.Sprintf(" %d", weight)
		}

		line += "\n"

		w.WriteString(line)
		w.Flush()
	}

	return nil
}
