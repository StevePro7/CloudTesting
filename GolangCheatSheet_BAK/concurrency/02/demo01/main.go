package main

import (
	"fmt"
	"github.com/nikandfor/goid"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}

func main() {

	for i :=0; i < 2; i++ {
		id := rnd.Intn(10) + 1
		fmt.Printf("[%d] ID:%d\n", goid.ID(), id)
		if b, ok := queryCache(id); ok {
			fmt.Printf("[%d] from cache\n", goid.ID())
			fmt.Printf("[%d] Book '%v'\n", goid.ID(), b.Title)
			continue
		}

		if b, ok :=queryDatabase(id); ok {
			fmt.Printf("[%d] from database\n", goid.ID())
			fmt.Printf("[%d] Book '%v'\n", goid.ID(), b.Title)
			continue
		}

		fmt.Printf("[%d] Book not found id '%v'\n", goid.ID(), id)
		time.Sleep(150 * time.Millisecond)
	}

}