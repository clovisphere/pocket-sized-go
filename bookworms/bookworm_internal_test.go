package main

import (
	"testing"
)

var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{
		Author: "Margaret Atwood", Title: "Oryx and Crake",
	}
	theBellJar = Book{
		Author: "Sylvia Plath", Title: "The Bell Jar",
	}
	janeEyre = Book{
		Author: "Charlotte Brontë", Title: "Jane Eyre",
	}
)

func Example_main() {
	main()
	// Output:
	// Here are the common books:
	//  - The Handmaid's Tale by Margaret Atwood
}

func TestBooksCount(t *testing.T) {
	tests := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{handmaidsTale, janeEyre, oryxAndCrake}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		"no bookworms": {input: []Bookworm{}, want: map[Book]uint{}},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, handmaidsTale, theBellJar}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    2,
			},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
			},
			want: map[Book]uint{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("expected %+v, got %+v", tc.want, got)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tests := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{handmaidsTale, janeEyre}},
			},
			want: []Book{handmaidsTale},
		},
		"three bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, janeEyre, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{handmaidsTale, janeEyre, oryxAndCrake}},
				{Name: "Charles", Books: []Book{handmaidsTale, janeEyre, oryxAndCrake}},
			},
			want: []Book{handmaidsTale, janeEyre, oryxAndCrake},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			if !equalBooks(t, got, tc.want) {
				t.Fatalf("expected %+v, got %+v", tc.want, got)
			}
		})
	}
}

func TestLoadWorms_Success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{handmaidsTale, janeEyre, oryxAndCrake}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got one - %s", err.Error())
			}

			if err == nil && tc.wantErr {
				t.Fatal("expected an error, go none")
			}

			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("expected %+v, got %+v", tc.want, got)
			}
		})
	}
}

// equalBookworms reports whether two slices of Bookworm are equal,
// ignoring order. A Bookworm is considered equal if both its Name
// matches and its list of Books matches (also ignoring order).
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	// seen keeps track of which elements in target have already been matched,
	// so that duplicates are handled correctly and we don't match the same
	// element more than once.
	seen := make([]bool, len(target))

	for _, bw0 := range bookworms {
		// found tracks whether bw0 matches any unseen element in target.
		found := false

		for j, bw1 := range target {
			if seen[j] {
				continue
			}
			if bw0.Name == bw1.Name && equalBooks(t, bw0.Books, bw1.Books) {
				seen[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// equalBooks reports whether two slices of Book are equal,
// ignoring order. Books are compared by value equality.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		return false
	}

	// seen keeps track of which elements in target have already been matched,
	// so that duplicates are handled correctly and we don't match the same
	// element more than once.
	seen := make([]bool, len(target))

	for _, b0 := range books {
		// found tracks whether b0 matches any unseen element in target.
		found := false

		for j, b1 := range target {
			if seen[j] {
				continue
			}
			if b0 == b1 {
				seen[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// equalBooksCount is a helper to test the equality
// of two maps of books count.
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}
