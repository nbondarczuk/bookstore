package inmemdb

import (
	"fmt"
	"sync"

	"bookstore/model"
)

// container of books guarded by a mutex
type InmemdbStore struct {
	m     sync.Mutex
	books map[string]model.Book
}

var store InmemdbStore

type InmemdbRepository struct{}

// Close unlocks access to possibly nonindempotent operation on books repository
func (InmemdbRepository) Close() {}

// GetBooks obtains all books
func (InmemdbRepository) GetBooks() []model.Book {
	store.m.Lock()
	defer store.m.Unlock()

	var books = make([]model.Book, len(store.books), len(store.books))
	for _, b := range store.books {
		books = append(books, b)
	}

	return books
}

// FindBookByID obtains a book by unique id provided
func (InmemdbRepository) FindBookByID(id string) (b model.Book, err error) {
	var ok bool

	store.m.Lock()
	defer store.m.Unlock()
	if b, ok = store.books[id]; ok {
		return b, nil
	}

	return b, fmt.Errorf("can not find book")
}

// AddBook adds new book to the repository
func (InmemdbRepository) AddBook(b model.Book) {
	store.m.Lock()
	defer store.m.Unlock()
	store.books[b.ID] = b
}
