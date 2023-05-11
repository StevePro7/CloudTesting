package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func foo() int {
	return 9
}

func main() {
	addr := ":8081"
	log.Printf("Starting server on %s\n", addr)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		bar := foo()
		msg := fmt.Sprintf("Hello foo '%d' at %s!\n", bar, now.Format("2006-01-02 15:04:05"))
		_, _ = io.WriteString(w, msg)
	})
	log.Fatal(http.ListenAndServe(addr, handlerFunc))
}
