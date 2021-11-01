package store

import (
	"sync"

	"github.com/emlog/goexample/store/factory"
)

// internal/store/memstore.go

package store

import (
mystore "bookstore/store"
factory "bookstore/store/factory"
"sync"
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
