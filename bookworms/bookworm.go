package main

import (
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

// Recommendations

type set map[Book]struct{}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)

	// Register all books on everyone's shelf
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooks := listOtherBooksOnShelf(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooks)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

func listOtherBooksOnShelf(i int, books []Book) []Book {
}
