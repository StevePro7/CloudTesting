package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num := 11
		fmt.Fprintf(w, "Hello, World='%d'\n", num)
	})

	http.ListenAndServe(":80", nil)
}
