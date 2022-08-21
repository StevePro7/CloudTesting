package main

import (
	"concurrencytesting/book"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	m.Lock()
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return book.Book{}, false
}

func main() {

	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	cacheCh := make(chan book.Book)
	dbCh := make(chan book.Book)

	wg.Add(2)
	for i := 0; i < 2; i++ {
		id := rnd.Intn(10) + 1

		// Send only channel
		go func(id, int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)

		// Send only channel

	}
}
