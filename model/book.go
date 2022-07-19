package model

// book represents adata bout a book
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `jsob:"isbn"`
	pages  uint
}

// Book simple container
type Books []Book

// Bookstore complex container
type Bookstore struct {
	books map[string]Book
}
