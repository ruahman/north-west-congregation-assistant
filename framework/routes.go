package framework

import (
	"net/http"
)

func Routes(mux *http.ServeMux) {
	mux.HandleFunc("/territory/", territoryHandler)
}
