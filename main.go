package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, we are here on path %s\n -- ", r.URL.Path)
		fmt.Fprintf(w, "Hello i received it %s\n from parameters", r.URL.Query())
	})

	http.ListenAndServe(":3333", nil)
}
