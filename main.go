package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":4000", nil)
}
