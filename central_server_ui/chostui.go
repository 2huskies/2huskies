package main

import (
	"17bit/calc/httputil"
	"flag"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	bind    string
	logFile string
)

var host *Host

type App struct {
	IndexHandler *IndexHandler
}

func (a *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = httputil.ShiftPath(req.URL.Path)
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
	flag.StringVar(&logFile, "l", "", "`файл` журнала")
	flag.Parse()

	var err error
	if logFile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		output, err := filepath.Abs(logFile)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		fo, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("open %s: %s\n", output, err)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				log.Fatalf("close %s\n", err)
			}
		}()
		logger = log.New(fo, "", log.LstdFlags)
	}
	a := &App{
		IndexHandler: new(IndexHandler),
	}
	logger.Printf("listening on: %s", bind)
	mime.AddExtensionType(".js", "text/javascript")
	err = http.ListenAndServe(bind, a)
	if err != nil {
		logger.Fatalf("Listen error: %s\n", err)
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
