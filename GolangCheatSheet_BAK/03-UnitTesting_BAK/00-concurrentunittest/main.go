package main

import (
	"fmt"
	"unittesting/foo"
)

func main() {
	fmt.Println("hello1")

	//h := "name"
	//foo.Track(h)
	//a := foo.Testing(h)
	//if !a {
	//	fmt.Println("not found")
	//}

	name := "name"
	hashtags := []string{name, "surname"}
	foo.SetHashtags(hashtags)
	foo.Untrack(name)
	a := foo.Testing(name)
	if !a {
		fmt.Println("not found")
	}
}
