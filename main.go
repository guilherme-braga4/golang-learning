package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Instance of http server
	r := mux.NewRouter()

	// Setting Routes
	routes(r)

	http.ListenAndServe(":3000", r)
}

func routes(r *mux.Router) {
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Its seems to be ok!"))
	})
	http.Handle("/", r)
}

func sendHello() {
	fmt.Println("Hello my friend")
}
