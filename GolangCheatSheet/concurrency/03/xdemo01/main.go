package main

import (
	"fmt"
)

var cache = map[int]Book{}
//var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fmt.Println("hello")
	fmt.Printf("%v", cache)
}