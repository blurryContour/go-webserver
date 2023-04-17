package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetBody(w http.ResponseWriter, r *http.Request) {
	log.Printf("\t%s\t%s\n", r.Host, r.RequestURI)

	body_bytes, err := io.ReadAll(r.Body)
	body := string(body_bytes)
	if err != nil {
		log.Printf("Error in reading body: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content := fmt.Sprintf("%s\n----\n\n%v", "Body", body)
	w.Write([]byte(content))
}
