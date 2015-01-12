package router

import (
	"github.com/gorilla/mux"
	"go/build"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	StaticDir = filepath.Join(defaultBase("github.com/ansible-in/prototype-anarcher"), "static")
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
