package models

type Edge struct {
	Source Node
	Target Node
	Params map[string]interface{}
}
