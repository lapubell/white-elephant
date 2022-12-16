package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type server struct {
	port           string
	clients        map[*websocket.Conn]Person
	messageChannel chan Message
}

func NewServer(port string) (*server, error) {
	s := server{
		port:           port,
		messageChannel: make(chan Message),
		clients:        map[*websocket.Conn]Person{},
	}

	// set up routes
	routes := map[string]func() http.HandlerFunc{
		"/ws": s.handleWebSocket,
		"/":   s.handleIndex,
	}
	for route, handler := range routes {
		http.HandleFunc(route, handler())
	}

	// start the web server message worker
	go s.broadcastPeople()
	go s.broadcastAvailableGifts()
	go s.handleMessages()

	return &s, nil
}

func (s *server) Serve() error {
	return http.ListenAndServe(":"+s.port, nil)
}
