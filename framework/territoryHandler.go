package framework

import (
	"net/http"
	"strconv"
)

func territoryHandler(w http.ResponseWriter, req *http.Request) {
	res := make(map[string]string)

	if req.URL.Path == "/territory/" {
		if req.Method == "GET" {
			res["message"] = "GET territory"
		} else if req.Method == "POST" {
			res["message"] = "POST territory"
		} else if req.Method == "PUT" {
			res["message"] = "PUT territory"
		} else if req.Method == "DELETE" {
			res["message"] = "DELETE territory"
		}
	} else {
		pathParts := getPathParts(req)
		if len(pathParts) < 2 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		res["message"] = "GET territory with id=" + strconv.Itoa(id)
	}

	renderJSON(w, res)
}
