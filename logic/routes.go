package logic

import (
	"net/http"
)

func Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("got path\n"))
	})

	mux.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_, _ = w.Write([]byte("handling task with id=" + id + "\n"))
	})

	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Query().Get("path")
		_, _ = w.Write([]byte("handling file with path=" + path + "\n"))
	})
}
