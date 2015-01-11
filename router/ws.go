package router

import "github.com/gorilla/mux"

const (
    WS = "ws"
)

func WebSocket() *mux.Router {
    m := mux.NewRouter()
    m.PathPrefix("/ws").Name(WS)
    return m
}

