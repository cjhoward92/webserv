package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	router := NewRouter()
	router.AddRoute(Route{Name: "Home", Path: "/", Handler: handler})
	router.Bind()

	http.ListenAndServe(":8080", nil)
}
