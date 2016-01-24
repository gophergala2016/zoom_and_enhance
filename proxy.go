package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"code.google.com/p/go-uuid/uuid"
	"golang.org/x/net/http2"
)

// Generate a unique request ID so that incoming requests can be associated with
// responding requests to origin server.
func genID() string {
	id := uuid.NewRandom()
	log.Println("UUID:", id.String())

	return id.String()
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func setHandlers(r *http.ServeMux) {
	r.HandleFunc("/", infoHandler)
	r.HandleFunc("/info", infoHandler)
}

func main() {
	var srv http.Server

	// Set Port
	srv.Addr = "localhost:8080"
	if httpAddr := os.Getenv("Address"); httpAddr != "" {
		srv.Addr = httpAddr
	}

	// Mux setup
	router := http.NewServeMux()

	// Set mux
	srv.Handler = router

	// Set handlers
	setHandlers(router)

	err := http2.ConfigureServer(&srv, &http2.Server{})
	if err != nil {
		log.Println("ConfigureServer Error:", err)
	}

	log.Printf("Listening on %s", srv.Addr)
	srv.ListenAndServeTLS("example.com.crt", "example.com.key")
}
