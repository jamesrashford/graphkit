package webui

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

//go:embed static/*
var content embed.FS

//go:embed templates/*
var templates embed.FS
var tmpl = template.Must(template.ParseFS(templates, "templates/*.html"))

func StartServer(addr string) {
	static := http.FileServer(http.FS(content))
	http.Handle("/static/", http.StripPrefix("/", static))

	http.HandleFunc("/", index)
	http.HandleFunc("/graph.json", graphEndpoint)

	// Step 3: Start the HTTP server
	log.Println("Serving on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Graphkit WebUI",
	}

	// Render the template with dynamic data
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Template rendering error:", err)
	}
}

func graphEndpoint(w http.ResponseWriter, r *http.Request) {
	gio1 := io.NewEdgeListIO("", "", true)
	var rw1 io.GraphIO = gio1

	file, err := os.Open("examples/lollipop/graph.edgelist")
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	graph, err := rw1.ReadGraph(file)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	gio2 := io.NewJSONIO()
	var rw2 io.GraphIO = gio2

	buf := new(bytes.Buffer)

	err = rw2.WriteGraph(graph, buf)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	data := models.JSONGraph{}
	json.Unmarshal(buf.Bytes(), &data)

	json.NewEncoder(w).Encode(data)
}
