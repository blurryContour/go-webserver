package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blurryContour/go-webserver/api"
	"github.com/blurryContour/go-webserver/args"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	// t := time.Now().Format(time.DateTime)
	log.Printf("\t%s\t%s\n", r.Host, r.RequestURI)

	content := fmt.Sprintf("%s\n", "Home Page!")
	w.Write([]byte(content))
}

func main() {
	args := args.ParseArgs(os.Args)

	mux := http.NewServeMux()

	// Add function handlers
	mux.HandleFunc("/", getHome)
	mux.HandleFunc("/body", api.GetBody)
	mux.HandleFunc("/form", api.GetForm)
	mux.HandleFunc("/formall", api.GetFormAll)

	// Start server
	addr := fmt.Sprintf(":%d", args.Port)
	fmt.Printf("\nStarting server on %s ...\n", addr)
	err := http.ListenAndServe(addr, mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed!\n")
	} else if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
