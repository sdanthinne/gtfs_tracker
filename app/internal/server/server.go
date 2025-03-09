package server

import (
	"log"
	"net/http"
)

func Create_server() *http.Server {
	log.Printf("Creating server")
	s := &http.Server{
		Addr: "8000",
	}
	return s
}

func Start_server(s *http.Server) {
	log.Printf("Starting server")
	s.ListenAndServe()
}
