package api

import (
	"fmt"
	"log"
	"net/http"
)

func GetForm(w http.ResponseWriter, r *http.Request) {
	log.Printf("\t%s\t%s\n", r.Host, r.RequestURI)

	queryKey := "color"
	queryValue := r.PostFormValue(queryKey)
	if queryValue == "" {
		queryValue = "none"
	}
	formData := map[string]string{queryKey: queryValue}

	content := fmt.Sprintf("%s\n---------\n%v", "Form Data", formData)
	w.Write([]byte(content))
}

func GetFormAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("\t%s\t%s\n", r.Host, r.RequestURI)

	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	formData := r.Form

	content := fmt.Sprintf("%s\n---------\n%v", "Form Data", formData)
	w.Write([]byte(content))
}
