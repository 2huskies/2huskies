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
	IndexHandler            *IndexHandler
	LoginHandler            *LoginHandler
	AbiturientHandler       *AbiturientHandler
	AbiturientScoresHandler *AbiturientScoresHandler
	SubjectsHandler         *SubjectsHandler
	SpecialtiesHandler      *SpecialtiesHandler
	UniversitiesHandler     *UniversitiesHandler
}

var api *ApiClient
var apiURL string

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
	case "login":
		a.LoginHandler.ServeHTTP(res, req)
		return
	case "abiturient":
		a.AbiturientHandler.ServeHTTP(res, req)
		return
	case "abiturient_scores":
		a.AbiturientScoresHandler.ServeHTTP(res, req)
		return
	case "subjects":
		a.SubjectsHandler.ServeHTTP(res, req)
		return
	case "specialties":
		a.SpecialtiesHandler.ServeHTTP(res, req)
		return
	case "universities":
		a.UniversitiesHandler.ServeHTTP(res, req)
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
	flag.StringVar(&apiURL, "api", "http://localhost:3081", "`URL` сервиса REST API")
	flag.Parse()

	var err error
	api, err = newApiClient(apiURL)
	if err != nil {
		log.Fatalf("%s", err)
	}
	a := &App{
		IndexHandler: new(IndexHandler),
		LoginHandler: new(LoginHandler),
	}
	initCache()
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
