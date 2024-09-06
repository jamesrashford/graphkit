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

	headerParams := make(map[string]bool)
	edges := graph.GetEdges()
	for _, e := range edges {
		for k, _ := range e.Params {
			headerParams[k] = true
		}
	}

	delimiter := ","

	headerIdx := make(map[int]string)
	header := "source" + delimiter + "target"
	i := 0
	for k, _ := range headerParams {
		headerIdx[i] = k
		header += fmt.Sprintf("%s%s", delimiter, k)
		i += 1
	}

	fmt.Println(header)

	for _, e := range edges {
		row := fmt.Sprintf("%v%s%v", e.Source.ID, delimiter, e.Target.ID)
		for i := 0; i < len(e.Params); i++ {
			v := e.Params[headerIdx[i]]
			row += fmt.Sprintf("%s%v", delimiter, v)
		}
		fmt.Println(row)
	}
}
