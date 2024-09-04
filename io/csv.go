package io

import (
	"encoding/csv"
	"io"

	"github.com/jamesrashford/graphkit/models"
)

type CSVIO struct {
	Comments  string
	Delimiter string
	SourceCol string
	TargetCol string
	Directed  bool
	Data      bool
}

func NewCSVIO(comment, delimiter string, source string, target string, directed bool, data bool) *CSVIO {
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
		Data:      data,
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

		// Check if other cols
		if csvio.Data {
			if len(headerMap) > 2 {
				params = make(map[string]interface{})

				for k, v := range headerMap {
					if k == csvio.SourceCol || k == csvio.TargetCol {
						continue
					}

					params[k] = record[v]
				}
			}
		}

		graph.AddEdge(record[headerMap[csvio.SourceCol]], record[headerMap[csvio.TargetCol]], params)
	}

	return graph, nil
}

func (csvio *CSVIO) WriteGraph(graph *models.Graph, writer io.Writer) error {
	header := make(map[string]int)
	header["source"] = 0
	header["target"] = 1

	edges := graph.GetEdges()
	for _, e := edges {
		i := 2
		for k, _ := range e.Params {
			header[k] = i
			i += 1
		}
	}

	return nil
}
