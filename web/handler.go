package web

import (
	"bitbucket.org/anarcher/ansible-prototype/router"
	"github.com/gorilla/mux"
)

func Handler() *mux.Router {
	m := router.WEB()
	return m
}

