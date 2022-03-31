package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/google/uuid"
)

var (
	//go:embed frontend/index.html
	html string
)

func handleMessages(messages []api.Message) {
	if len(messages) < 1 {
		return
	}

	for _, message := range messages {
		detail := fmt.Sprintf(
			"\n%s\n    %s:%d:%d\n      %d | %s",
			message.Text,
			message.Location.File,
			message.Location.Line,
			message.Location.Column,
			message.Location.Line,
			message.Location.LineText,
		)

		log.Println(detail)
	}

	os.Exit(1)
}

func bundle() template.JS {
	bundled := api.Build(api.BuildOptions{
		EntryPoints: []string{"frontend/App.jsx"},
		Bundle:      true,
	})
	handleMessages(bundled.Errors)

	return template.JS(bundled.OutputFiles[0].Contents) // nolint
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	tmpl, _ := template.New("html").Parse(html) // nolint

	data := struct{ JS template.JS }{JS: bundle()}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data) // nolint
	})

	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) { // nolint
		source := map[string]interface{}{
			"id":        uuid.New(),
			"name":      "Polo",
			"email":     "meetup_gdl@pitakill.net",
			"languages": []string{"Go", "JS", "bash"},
		}

		data, _ := json.Marshal(source) // nolint

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(data) // nolint
	})

	http.ListenAndServe(port, nil) // nolint
}
