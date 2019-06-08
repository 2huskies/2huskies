package main

import (
	"io/ioutil"
	"net/http"
)

type IndexHandler struct {
}

func (h *IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadFile("html/index.html")
	_, _ = res.Write(buf)
}
