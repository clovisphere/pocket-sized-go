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

// set is a collection of unique Books.
type set map[Book]struct{}

// NewSet returns a set containing the given books.
func NewSet(books ...Book) set {
	s := make(set)
	s.Add(books...)
	return s
}

// Add inserts the given books into the set.
func (s set) Add(books ...Book) {
	for _, b := range books {
		s[b] = struct{}{}
	}
}

// Contains reports whether b is in the set.
func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
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

// findCommonBooks returns books that are on more than one bookworm's shelf.
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

// recommendOtherBooks returns a slice of Bookworm where each Bookworm contains
// the same Name as the input, but Books replaced with recommendations. A
// recommendation for a Bookworm consists of books found on other Bookworms'
// shelves that the given Bookworm has not read. Each recommended book appears
// at most once per Bookworm.
func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	var bw []Bookworm

	for _, reader := range bookworms {
		seen := NewSet(reader.Books...)
		var unread []Book

		for _, peer := range bookworms {
			// Skip recommending from oneself
			if reader.Name == peer.Name {
				continue
			}

			for _, book := range peer.Books {
				if seen.Contains(book) {
					continue
				}
				unread = append(unread, book)
				seen.Add(book)
			}
		}

		bw = append(bw, Bookworm{
			Name:  reader.Name,
			Books: unread,
		})
	}

	return bw
}

// sortBooks sorts the books by Author and then Title
// in alphabetical order.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}
