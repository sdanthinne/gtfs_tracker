package server

import (
    "net/http"
    "log"
)

func Create_server() Server {
    log.Printf("Creating server")
    s := &http.Server{
    	Addr: "8000",
    }
    return s
}

go func Start_server( s Server ) {
    log.Printf("Starting server")
    s.ListenAndServe()
}
