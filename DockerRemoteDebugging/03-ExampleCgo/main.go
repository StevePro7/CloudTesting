package main

import (
	"fmt"
	"net/http"
)

// #include "foo.h"
import "C"

func Bar() int {
	//bar := int(C.Cbar())
	bar := 4
	return bar
}

func main() {
	num1 := Bar()
	fmt.Printf("starting...:'%d'\n", num1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num2 := Bar()
		fmt.Fprintf(w, "Hello, World='%d'\n", num2)
	})

	http.ListenAndServe(":80", nil)
}
