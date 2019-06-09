package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FacultiesHandler struct{}

func (h *FacultiesHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	result, err := api.getFaculties(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("get faculties: %s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(result)
}
