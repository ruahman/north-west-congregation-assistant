package framework

import (
	"encoding/json"
	"net/http"
	"strings"
)

func getPathParts(req *http.Request) []string {
	path := req.URL.Path
	path = strings.Trim(path, "/")
	return strings.Split(path, "/")
}

func renderJSON(w http.ResponseWriter, res interface{}) {
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
