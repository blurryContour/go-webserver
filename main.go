package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/blurryContour/go-webserver/api"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	// t := time.Now().Format(time.DateTime)
	log.Printf("\t%s\t%s\n", r.Host, r.RequestURI)

	content := fmt.Sprintf("%s\n", "Home Page!")
	w.Write([]byte(content))
}

func main() {
	mux := http.NewServeMux()

	// Add function handlers
	mux.HandleFunc("/", getHome)
	mux.HandleFunc("/body", api.GetBody)
	mux.HandleFunc("/form", api.GetForm)
	mux.HandleFunc("/formall", api.GetFormAll)

	// Start server
	fmt.Printf("\nStarting server...\n")
	err := http.ListenAndServe(":80", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed!\n")
	} else if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
