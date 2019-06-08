package main

import (
	"flag"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"
)

var (
	bind    string
	logFile string
)

type App struct {
	IndexHandler *IndexHandler
}

func (a *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "":
		a.IndexHandler.ServeHTTP(res, req)
		return
	case "ping":
		res.Write([]byte("pong"))
		return
	case "html", "js", "css", "fonts":
		http.FileServer(http.Dir(head)).ServeHTTP(res, req)
		return
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func main() {
	flag.StringVar(&bind, "b", ":3080", "`адрес:порт` на котором слушать запросы")
	flag.Parse()

	var err error
	a := &App{
		IndexHandler: new(IndexHandler),
	}
	log.Printf("listening on: %s", bind)
	mime.AddExtensionType(".js", "text/javascript")
	err = http.ListenAndServe(bind, a)
	if err != nil {
		log.Fatalf("Listen error: %s\n", err)
	}
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
