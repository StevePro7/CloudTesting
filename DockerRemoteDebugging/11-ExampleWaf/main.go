package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testwebapi/server"
)

func main() {

	bind := ":8081"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", server.AdminHelloHandler).Methods("GET")
	gmux.HandleFunc("/test/artists.php", server.TestHandler).Methods("GET")
	log.Printf("starting smart reverse proxy on [%s]", bind)

	if err := http.ListenAndServe(bind, server.LimitMiddleware(gmux)); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
