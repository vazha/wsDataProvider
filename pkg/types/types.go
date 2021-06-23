package types

import (
	"fmt"
	"sync"
)

type Book struct {
	LastUpdated int64
	Pair string
	Asks []float64
	Bids []float64
}

// BookStore represent global map of books
type BookStore struct {
	*sync.RWMutex
	books map[string][]*Book // map[Exchange] []Book
	//Books map[string]map[string]Book
}

// NewBookStore creates new global instance
func NewBookStore() *BookStore {
	return &BookStore{
		books: make(map[string][]*Book),
		RWMutex: &sync.RWMutex{},
	}
}

// Get book by Exchange and Pair
func (o *BookStore) Get(exchange, pair string) (b *Book, err error) {
	o.RLock()
	if book, ok := o.books[exchange]; ok {
		o.RUnlock()
		for i := range book {
			if book[i].Pair == pair {
				return book[i], nil
			}
		}
		return nil, fmt.Errorf("no pair found")
	}
	o.RUnlock()
	return nil, fmt.Errorf("no exchange found")
}

// Set book by Exchange and Pair
func (o *BookStore) Set(exchange string, b *Book) {
	o.Lock()
	defer o.Unlock()
	if book, ok := o.books[exchange]; ok {
		for i := range book {
			if book[i].Pair == b.Pair {
				book[i] = b
				return
			}
		}
		o.books[exchange] = append(o.books[exchange], b)
		return
	}

	o.books[exchange] = []*Book{b}
	return
}