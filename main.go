package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jamesrashford/graphkit/models"
)

func main() {
	file, err := os.Open("examples/lollipop/graph_data.csv")
	if err != nil {
		panic(err)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	headerMap := make(map[string]int)

	graph := models.NewEmptyGraph(true)

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
				if k == "source" || k == "target" {
					continue
				}
				params[k] = record[v]
			}

		}
		graph.AddEdge(record[headerMap["source"]], record[headerMap["target"]], params)
	}

	edges := graph.GetEdges()
	for _, e := range edges {
		for k, _ := range e.Params {
			fmt.Println(k)
		}
	}
	fmt.Println(edges)
}
