package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testwebapi/foo"
	"time"
)

// https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code
func main() {
	addr := ":8081"
	log.Printf("Starting server on %s\n", addr)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		bar := foo.Bar()
		msg := fmt.Sprintf("Hello foo '%d' at %s!\n", bar, now.Format("2006-01-02 15:04:05"))
		_, _ = io.WriteString(w, msg)
	})
	log.Fatal(http.ListenAndServe(addr, handlerFunc))
}
