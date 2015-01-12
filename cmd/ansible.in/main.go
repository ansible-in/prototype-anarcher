package main

import (
	"flag"
	"github.com/ansible-in/prototype-anarcher"
	"github.com/ansible-in/prototype-anarcher/web"
	"github.com/ansible-in/prototype-anarcher/ws"
	"log"
	"net/http"
)

var (
	httpAddr = flag.String("http", ":5000", "HTTP service address")
)

func main() {
	flag.Parse()
	ansible.InitDB()

	m := http.NewServeMux()
	//m.Handle("/ws/",http.StripPrefix("/ws",ws.Handler()))
	m.Handle("/ws/", ws.Handler())
	m.Handle("/", web.Handler())

	log.Print("Listening on ", *httpAddr)
	err := http.ListenAndServe(*httpAddr, m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
