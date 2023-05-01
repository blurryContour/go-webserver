package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/blurryContour/go-webserver/core"
)

func _getMapFromRequest(r *http.Request, keys []string) map[string]string {
	data := make(map[string]string, len(keys))
	for _, k := range keys {
		data[k] = r.PostFormValue(k)
	}
	return data
}

func _compute(w http.ResponseWriter, r *http.Request, routine bool) {
	log.Printf("\t%s\t%s\n", r.Host, r.RequestURI)

	queryKey := "n"
	formData := _getMapFromRequest(r, []string{queryKey})
	n, err := strconv.Atoi(formData[queryKey])
	if err != nil {
		log.Fatalf("Invalid number: %s\n", formData[queryKey])
	}

	content := core.Compute(n, routine)
	w.Write([]byte(content))
}

func ComputeAsync(w http.ResponseWriter, r *http.Request) {
	_compute(w, r, true)
}
func ComputeSync(w http.ResponseWriter, r *http.Request) {
	_compute(w, r, false)
}
