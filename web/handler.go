package web

import (
	"github.com/ansible-in/prototype-anarcher/router"
	"github.com/gorilla/mux"
)

func Handler() *mux.Router {
	m := router.WEB()
	return m
}
