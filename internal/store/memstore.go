package store

import (
	"sync"

	mystore "goexample/store"
	factory "goexample/store/factory"

	"github.com/emlog/goexample/store/factory"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}
