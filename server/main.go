package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	publish = flag.String("publish", "public", "Directory to serve content as root")
	port    = flag.Int("port", 8080, "Port for the http server")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/_/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(*publish)))
	log.Printf("Listening on port %d", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}
