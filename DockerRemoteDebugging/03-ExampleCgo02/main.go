package main

import (
	"fmt"
	"net/http"

	"github.com/stevepro/remote-debug-example/foo"
)

func main() {
	num1 := foo.Bar()
	fmt.Printf("starting...:'%d'\n", num1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num2 := foo.Bar()
		fmt.Fprintf(w, "Hello, World='%d'\n", num2)
	})

	http.ListenAndServe(":80", nil)
}
