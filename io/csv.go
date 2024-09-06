package io

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/jamesrashford/graphkit/models"
)

type CSVIO struct {
	Comments  string
	Delimiter string
	SourceCol string
	TargetCol string
	Directed  bool
}

func NewCSVIO(comment, delimiter string, source string, target string, directed bool) *CSVIO {
	if comment == "" {
		comment = "#"
	}
	if delimiter == "" {
		delimiter = ","
	}

	if source == "" {
		source = "source"
	}

	if target == "" {
		target = "target"
	}

	csvio := CSVIO{
		Comments:  comment,
		Delimiter: delimiter,
		SourceCol: source,
		TargetCol: target,
		Directed:  directed,
	}
	return &csvio
}

func (csvio *CSVIO) ReadGraph(reader io.Reader) (*models.Graph, error) {
	graph := models.NewEmptyGraph(csvio.Directed)

	records, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return nil, err
	}

	headerMap := make(map[string]int)

	for i, record := range records {
		if i == 0 {
			for j, h := range record {
				headerMap[h] = j
			}
			continue
		}
		var params map[string]interface{}

		if len(headerMap) > 2 {
			params = make(map[string]interface{})

			for k, v := range headerMap {
				if k == csvio.SourceCol || k == csvio.TargetCol {
					continue
				}

				params[k] = record[v]
			}
		}

		graph.AddEdge(record[headerMap[csvio.SourceCol]], record[headerMap[csvio.TargetCol]], params)
	}

	return graph, nil
}

func (csvio *CSVIO) WriteGraph(graph *models.Graph, writer io.Writer) error {
	headerParams := make(map[string]bool)
	edges := graph.GetEdges()
	for _, e := range edges {
		for k, _ := range e.Params {
			headerParams[k] = true
		}
	}

	delimiter := csvio.Delimiter

	headerIdx := make(map[int]string)
	header := "source" + delimiter + "target"
	i := 0
	for k, _ := range headerParams {
		headerIdx[i] = k
		header += fmt.Sprintf("%s%s", delimiter, k)
		i += 1
	}
	header += "\n"

	w := bufio.NewWriter(writer)
	w.WriteString(header)
	w.Flush()

	for _, e := range edges {
		line := fmt.Sprintf("%v%s%v", e.Source.ID, delimiter, e.Target.ID)
		for i := 0; i < len(e.Params); i++ {
			v := e.Params[headerIdx[i]]
			line += fmt.Sprintf("%s%v", delimiter, v)
		}
		line += "\n"

		w.WriteString(line)
		w.Flush()
	}

	return nil
}
