package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id     string `json:"id"`     // 图书ISBN ID
	Name   string `json:"name"`   // 图书名称
	Author string `json:"author"` // 图书作者
	Press  string `json:"press"`  // 出版社
}

type Store interface {
	Create(*Book) error
	Get(string) (Book, error)
}
