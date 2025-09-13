package main

import (
	"encoding/json"
	"os"
	"sort"
)

// A Bookworm contains the list of books on a bookworn's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// byAuthor is a list of Book.
// Defining a custom type to implement sort.Interface
type byAuthor []Book

// Len implements sort.Interface by
// returning the length of the book per author.
func (b byAuthor) Len() int { return len(b) }

// Swap implements sort.Interface and swaps two books.
func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements sort.Interface and
// returns books sorted by Author then Title.
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// booksCount registers all the books and their occurrences
// from the bookworms shelves.
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// findCommonBooks returns books are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

// loadBookworms reads the file and returns the list of bookworms,
// and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// sortBooks sorts the books by Author and then Title
// in alphabetical order.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}
