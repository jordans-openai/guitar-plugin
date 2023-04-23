package server

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/dustmason/guitar-plugin/tablature"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed ai-plugin.json
var aiPluginJSON string

//go:embed openapi.yaml
var openapiYAML string

type Chord struct {
	Name string `json:"name"`
}

type ChordsRequest struct {
	Chords []Chord `json:"chords"`
}

type TabsResponse struct {
	Tabs []string `json:"tabs"`
}

type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	})
}

func ChordsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var chordsReq ChordsRequest
	err := json.NewDecoder(r.Body).Decode(&chordsReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tabs := make([]string, len(chordsReq.Chords))
	for i, cr := range chordsReq.Chords {
		tab, err := tablature.NewTablature(cr.Name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid chord name: %s", cr.Name), http.StatusBadRequest)
			return
		}
		tabs[i] = tab.String()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TabsResponse{Tabs: tabs})
}

func PluginManifestHandler(w http.ResponseWriter, r *http.Request) {
	replaced := replaceHostname(aiPluginJSON, r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(replaced))
}

func OpenAPISpecHandler(w http.ResponseWriter, r *http.Request) {
	replaced := replaceHostname(openapiYAML, r)
	w.Header().Set("Content-Type", "application/x-yaml")
	w.Write([]byte(replaced))
}

func CORSHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func replaceHostname(content string, r *http.Request) string {
	protocol := "http"
	if os.Getenv("DEV") == "" {
		protocol = "https"
	}
	return strings.Replace(content, "PLUGIN_HOSTNAME", fmt.Sprintf("%s://%s", protocol, r.Host), -1)
}
