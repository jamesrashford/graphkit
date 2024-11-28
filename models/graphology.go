package models

type GraphologyGraph struct {
	Attributes map[string]interface{} `json:"attributes"`
	Options    map[string]interface{} `json:"options"`
	Nodes      []struct {
		Key        string                 `json:"key"`
		Attributes map[string]interface{} `json:"attributes"`
	} `json:"nodes"`
	Edges []struct {
		Key        string                 `json:"key"`
		Source     string                 `json:"source"`
		Target     string                 `json:"target"`
		Attributes map[string]interface{} `json:"attributes"`
	} `json:"edges"`
}
