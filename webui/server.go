package webui

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/models"
)

//go:embed static/*
var content embed.FS

//go:embed templates/*
var templates embed.FS
var tmpl = template.Must(template.ParseFS(templates, "templates/*.html"))

type WebUI struct {
	Address string
	Graph   models.Graph
}

func NewWebUI(addr string, graph models.Graph) *WebUI {
	return &WebUI{
		Address: addr,
		Graph:   graph,
	}
}

func (ui *WebUI) StartServer() {
	static := http.FileServer(http.FS(content))
	http.Handle("/static/", http.StripPrefix("/", static))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/graph.json", func(w http.ResponseWriter, r *http.Request) {
		rw := io.NewJSONIO()
		buf := new(bytes.Buffer)

		err := rw.WriteGraph(&ui.Graph, buf)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		data := models.JSONGraph{}
		json.Unmarshal(buf.Bytes(), &data)

		json.NewEncoder(w).Encode(data)
	})

	log.Printf("Serving WebUI on %s...\n", ui.Address)
	err := http.ListenAndServe(ui.Address, logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
