package framework

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	Routes(mux)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
