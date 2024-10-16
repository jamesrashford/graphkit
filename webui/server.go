package webui

import (
	"embed"
	"html/template"
	"log"
	"net/http"
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
