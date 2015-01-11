package router

import (
    "github.com/gorilla/mux"
    "path/filepath"
    "net/http"
	"go/build"
    "os"
    "log"
)

var (
	StaticDir = filepath.Join(defaultBase("bitbucket.org/anarcher/ansible-prototype"), "static")
)

func WEB() *mux.Router {
    m := mux.NewRouter()
    m.PathPrefix("/").Handler(http.FileServer(http.Dir(StaticDir)))

    return m
}

func defaultBase(path string) string {
	p, err := build.Default.Import(path, "", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p.Dir, err = filepath.Rel(cwd, p.Dir)
	if err != nil {
		log.Fatal(err)
	}

	return p.Dir
}

