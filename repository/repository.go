package repository

import (
	"fmt"

	"bookstore/common/config"
	"bookstore/model"
	"bookstore/repository/inmemdb"
)

type Repository interface {
	Close()
	GetBooks() []model.Book
	FindBookByID(string) (model.Book, error)
	AddBook(model.Book)
}

// NewRepository is a factory function creating specific repository
// depending on the configuration
func NewRepository() (Repository, error) {
	if config.RepositoryType == "inmemdb" {
		return inmemdb.InmemdbRepository{}, nil
	}

	return nil, fmt.Errorf("invalid repository type")
}
