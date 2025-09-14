package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "path", "testdata/bookworms.json", "Path to the JSON file containing Bookworms data")
	flag.Parse()

	bookworms, err := loadBookworms(filePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	// pretty print `bookworms`
	// b, _ := json.MarshalIndent(bookworms, "", " ")
	// fmt.Println(string(b))

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("\n** Common books **")
	displayBooks(commonBooks)

	recommendedBooks := recommendOtherBooks(bookworms)

	fmt.Println("\n** Recommended books **")
	displayRecommendations(recommendedBooks)
}

// displayBooks prints out the titles and authors of a list of books.
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}
}

// displayRecommendations prints out book recommendations for each Bookworm.
func displayRecommendations(bookworms []Bookworm) {
	for _, bookworm := range bookworms {
		fmt.Printf("%s, we think you may also like:\n", bookworm.Name)
		displayBooks(bookworm.Books)
	}
}
