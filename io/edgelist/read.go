package edgelist

import (
	"bufio"
	"os"
	"strings"

	"github.com/jamesrashford/graphkit/models"
)

func ReadGraph(path string, directed bool) (*models.Graph, error) {
	graph := models.NewEmptyGraph(directed)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")

		// check for json data in idx 2. If so, load into params
		source := row[0]
		target := row[1]

		graph.AddEdge(source, target, nil)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}
