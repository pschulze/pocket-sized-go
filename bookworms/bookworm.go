package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	// Standard file opening
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Using a buffered reader instead of directly passing the file to json.Decoder
	// to reduce number of read syscalls, partiuclarly for larger files.
	// Buffer size can be adjusted; useful if size is (relatively) well known.
	buffedReader := bufio.NewReaderSize(f, 1024*1024) // 1MB buffer
	// buffio.Reader doesn't implement Closer, no need to defer
	decoder := json.NewDecoder(buffedReader)

	var bookworms []Bookworm
	err = decoder.Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	var commonBooks []Book
	for book, count := range booksCount(bookworms) {
		// TODO - best practices for uint len? Just return map[Book]int instead?
		if count == uint(len(bookworms)) {
			commonBooks = append(commonBooks, book)
		}
	}

	sortBooks(commonBooks)
	return commonBooks
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		uniqueBooks := make(map[Book]bool)
		for _, book := range bookworm.Books {
			// Don't count duplicate books from the same bookworm
			found := uniqueBooks[book]
			if !found {
				uniqueBooks[book] = true
				count[book]++
			}

		}
	}
	return count
}

// List of books wtih a custom type for defining sort interface.
// To implement sort interface, define Len, Less, and Swap methods.
type byAuthor []Book

func (b byAuthor) Len() int { return len(b) }

func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}

	return b[i].Title < b[j].Title
}

func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

// Alternative implementation using slices.SortFunc
func sortBooksSlicesSortFunc(books []Book) {
	slices.SortFunc(books, func(a, b Book) int {
		if a.Author != b.Author {
			return strings.Compare(a.Author, b.Author)
		}
		return strings.Compare(a.Title, b.Title)
	})
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
