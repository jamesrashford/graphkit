package models

type JSONGraph struct {
	Directed   bool `json:"directed"`
	Multigraph bool `json:"multigraph"`
	Graph      struct {
	} `json:"graph"`
	Nodes []struct {
		ID interface{} `json:"id"`
	} `json:"nodes"`
	Links []struct {
		Source interface{} `json:"source"`
		Target interface{} `json:"target"`
	} `json:"links"`
}
