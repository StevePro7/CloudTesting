package main

import (
	"fmt"
	"github.com/stevepro/remote-debug-example/foo"
)

func main() {
	bar := foo.Bar()
	fmt.Printf("Bar='%d'\n", bar)
}
