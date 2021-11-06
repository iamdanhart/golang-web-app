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

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

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
