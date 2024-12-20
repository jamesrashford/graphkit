package cmd

type FormatType int64

const (
	EdgeListType FormatType = iota
	CSVType
	GraphologyJSONType
	NodeLinkJSONType
)

var TypeName = map[FormatType]string{
	EdgeListType:       "edgelist",
	CSVType:            "csv",
	GraphologyJSONType: "graphology",
	NodeLinkJSONType:   "json",
}

var TypeDescription = map[FormatType]string{
	EdgeListType:       "A simple two-column edgelist txt file",
	CSVType:            "A CSV file as a two-column edgelist file",
	GraphologyJSONType: "A Graphology as a JSON file",
	NodeLinkJSONType:   "A JSON file in the node-link format",
}
