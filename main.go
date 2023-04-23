package main

import (
	"fmt"
	"github.com/dustmason/guitar-plugin/server"
	"log"
	"net/http"
)

func main() {
	http.Handle("/chords", server.ApplyMiddlewares(http.HandlerFunc(server.ChordsHandler), server.LoggingHandler, server.CORSHandler))
	http.Handle("/.well-known/ai-plugin.json", server.ApplyMiddlewares(http.HandlerFunc(server.PluginManifestHandler), server.LoggingHandler, server.CORSHandler))
	http.Handle("/openapi.yaml", server.ApplyMiddlewares(http.HandlerFunc(server.OpenAPISpecHandler), server.LoggingHandler, server.CORSHandler))
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
