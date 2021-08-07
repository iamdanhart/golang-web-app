package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	r := newRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Failed to start server, %s\n", err.Error())
		os.Exit(1)
	}
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, _ *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		fmt.Printf("Issue writing response: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}