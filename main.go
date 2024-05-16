package main

import (
	"net/http"
)

func main() {
	// Instance of http server
	mux := http.NewServeMux()
	http.ListenAndServe(":3000", mux)
}
