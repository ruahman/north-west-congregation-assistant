package logic

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	Routes(mux)

	err := http.ListenAndServe("localhost:8090", mux)
	if err != nil {
		log.Fatal(err)
	}
}
