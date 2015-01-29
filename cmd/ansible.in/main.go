package main

import (
	"flag"
	"github.com/ansible-in/prototype-anarcher"
	"github.com/ansible-in/prototype-anarcher/web"
	"github.com/ansible-in/prototype-anarcher/ws"
	"log"
	"net/http"
	"os"
)

var (
	host = flag.String("host", "", "HTTP service address")
	port = flag.String("port", "5000", "HTTP service port")
)

func main() {
	flag.Parse()
	ansible.InitDB()

	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}
	httpAddr := *host + ":" + *port

	m := http.NewServeMux()
	//m.Handle("/ws/",http.StripPrefix("/ws",ws.Handler()))
	m.Handle("/ws/", ws.Handler())
	m.Handle("/", web.Handler())

	log.Print("Listening on ", httpAddr)
	err := http.ListenAndServe(httpAddr, m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
