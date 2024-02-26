package framework

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Server() {
	mux := http.NewServeMux()

	Routes(mux)
	port := os.Getenv("PORT")
	fmt.Println("Server running on port", port)
	err := http.ListenAndServe("localhost:"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
