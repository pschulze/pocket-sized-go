package main

import (
	"slices"
)

// Maps a book to a collection of books read by other bookworms who have read that book.
type bookRecommendations map[Book]bookCollection

// A collection of books as a set for easy lookup and cheap storage.
type bookCollection map[Book]struct{}

// create a bookRecommendations
// Loop over all bookworms and their books
// 	 For each book, register all other books of that bookworm to that Book in the bookRecommendations map
// Loop over all bookworms and their books again
// 	 For each book, add all of it's recommendations to a new list of books for that bookworm

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	recommendations := buildBookRecommendations(bookworms)
	bookwormRecommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		bookRecs := make([]Book, 0)

		// Collect recommendations for all books of this bookworm
		// Make sure to remove duplicates and anything already on their shelf
		for _, book := range bookworm.Books {
			if recs, exists := recommendations[book]; exists {
				for b := range recs {
					if !slices.Contains(bookworm.Books, b) && !slices.Contains(bookRecs, b) {
						bookRecs = append(bookRecs, b)
					}
				}
			}
		}

		bookwormRecommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: bookRecs,
		}
	}

	return bookwormRecommendations
}

func buildBookRecommendations(bookworms []Bookworm) bookRecommendations {
	bookRecs := make(bookRecommendations)
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooks := slices.Delete(slices.Clone(bookworm.Books), i, i+1)
			// Register otherBooks as recommendations for book
			if _, exists := bookRecs[book]; !exists && len(otherBooks) > 0 {
				bookRecs[book] = make(bookCollection)
			}
			for _, otherBook := range otherBooks {
				bookRecs[book][otherBook] = struct{}{}
			}
		}
	}

	return bookRecs
}
