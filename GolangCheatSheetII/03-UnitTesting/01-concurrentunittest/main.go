package main

import (
	"fmt"
	"unittesting/foo"
)

func main() {
	fmt.Println("hello")

	y := foo.Splat()
	fmt.Printf("%d\n", y)
}
