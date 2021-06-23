package types

import "sync"

type Book struct {
	LastUpdated int64
	Pair string
	Asks []float64
	Bids []float64
}

// AllBooks represent global map of books
type AllBooks struct {
	sync.RWMutex
	//Books map[string]map[string]Book
	Books map[string]*Book // map[Exchange]Book
}
