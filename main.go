package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from foo!")

}

func main() {

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/", hello)
	server.ListenAndServe()

}
