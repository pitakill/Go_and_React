package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	// Over simplified to expose an idea not an architecture
	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		source := map[string]interface{}{
			"id":        uuid.New(),
			"name":      "Polo",
			"email":     "meetup_gdl@pitakill.net",
			"languages": []string{"Go", "JS", "bash"},
		}

		data, _ := json.Marshal(source)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(data) // nolint
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
