package models

type Graph struct {
	Nodes    map[string]map[string]interface{}
	Edges    map[string]map[string]map[string]interface{}
	Params   map[string]interface{}
	Directed bool
}

func NewEmptyGraph(directed bool) *Graph {
	return NewGraph(nil, nil, nil, directed)
}

func NewGraph(nodes map[string]map[string]interface{}, edges map[string]map[string]map[string]interface{}, params map[string]interface{}, directed bool) *Graph {
	graph := Graph{Nodes: nodes, Edges: edges, Params: params, Directed: directed}
	if graph.Nodes == nil {
		graph.Nodes = make(map[string]map[string]interface{})
	}

	if graph.Edges == nil {
		graph.Edges = make(map[string]map[string]map[string]interface{})
	}

	if graph.Params == nil {
		graph.Params = make(map[string]interface{})
	}

	return &graph
}

func (g Graph) HasNode(node string) bool {
	_, ok := g.Nodes[node]
	return ok
}

func (g *Graph) AddNode(node string) {
	if !g.HasNode(node) {
		g.Nodes[node] = make(map[string]interface{})
		g.Nodes[node]["ID"] = node
		g.Nodes[node]["Label"] = node
	}
}

func (g Graph) HasEdge(from, to string) bool {
	_, ok := g.Edges[from][to]
	return ok
}

func (g Graph) AddEdge(from, to string) {
	// Check if from and to are nodes. If not, add.
	if !g.HasNode(from) {
		g.AddNode(from)
	}
	if !g.HasNode(to) {
		g.AddNode(to)
	}

	if !g.HasEdge(from, to) {
		_, ok := g.Edges[from]
		if !ok {
			g.Edges[from] = make(map[string]map[string]interface{})
		}

		g.Edges[from][to] = make(map[string]interface{})
	}
}

func (g Graph) GetNodes() []Node {
	var nodes []Node
	for id, params := range g.Nodes {
		node := Node{ID: id, Label: id, Params: params}
		nodes = append(nodes, node)
	}

	return nodes
}

func (g Graph) GetEdges() []Edge {
	var edges []Edge

	for s, ts := range g.Edges {
		for t, params := range ts {
			sp := g.Nodes[s]
			tp := g.Nodes[t]

			edge := Edge{
				Source: Node{
					ID:     sp["ID"].(string),
					Label:  sp["Label"].(string),
					Params: sp},
				Target: Node{
					ID:     tp["ID"].(string),
					Label:  tp["Label"].(string),
					Params: tp},
				Params: params}
			edges = append(edges, edge)
		}
	}

	return edges
}
