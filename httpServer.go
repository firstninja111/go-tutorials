package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL type: %T\n", r.URL)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}