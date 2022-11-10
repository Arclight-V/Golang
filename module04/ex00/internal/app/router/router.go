package router

import (
	"io"
	"log"
	"net"
	"time"
)

type Router struct {
	listener net.Listener
}

func NewRouter(network, adress string) *Router {
	listener, err := net.Listen(network, adress)
	if err != nil {
		log.Fatal(err)
	}
	return &Router{listener: listener}
}

func (r *Router) handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (r *Router) Start() {
	for {
		conn, err := Router.listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
	}
}
