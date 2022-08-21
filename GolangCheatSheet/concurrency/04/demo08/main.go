package main

import (
	"concurrencytesting/book"
	"fmt"
	"github.com/nikandfor/goid"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan book.Book)
	dbCh := make(chan book.Book)

	for i := 0; i < 2; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)

		// Send only channel
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)

		// Send only channel
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// Receive only channels
		go func(cacheCh, dbCh <-chan book.Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh

			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}

	log.Printf("[%d] running.", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE", goid.ID())
}

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			return b, true
		}
	}

	return book.Book{}, false
}
