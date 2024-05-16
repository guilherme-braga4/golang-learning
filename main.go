package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Instance of http server
	r := mux.NewRouter()

	// Setting Routes
	routes(r)

	// Connect the server to the TCP connection
	http.ListenAndServe(":3000", r)
}

func routes(r *mux.Router) {
	r.HandleFunc("/hello", sendHello)
}

func sendHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Its seems to be ok!"))
}
