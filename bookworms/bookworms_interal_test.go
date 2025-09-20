package main

import (
	"reflect"
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

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
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
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if err == nil && tc.wantErr {
				t.Fatalf("expected an error, got none")
			}

			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}

			// Note: Can also use reflect.DeepEqual here instead of helper methods.
			// It's less performant, but less code to write.
			// Per Learn Go with Pocket-Sized Projects, not recommended for
			// production use.
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  map[Book]uint
	}

	testCases := map[string]testCase{
		"bookworms have books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"one bookworm has no books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{handmaidsTale: 1, theBellJar: 1},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)

			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  []Book
	}

	testCases := map[string]testCase{
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{handmaidsTale, theBellJar}},
			},
			want: []Book{handmaidsTale},
		},
		"many common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{handmaidsTale, oryxAndCrake}},
			},
			want: []Book{oryxAndCrake, handmaidsTale},
		},
		"no common books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{oryxAndCrake}},
				{Name: "Peggy", Books: []Book{theBellJar}},
			},
			want: []Book{},
		},
		"a bookworm has read no books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{oryxAndCrake}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: []Book{},
		},
		"neither bookworm has read any books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: []Book{},
		},
		"one bookworm has read two copies of the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, handmaidsTale}},
				{Name: "Peggy", Books: []Book{theBellJar}},
			},
			want: []Book{},
		},
		"three bookwoms have one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Bob", Books: []Book{handmaidsTale, theBellJar}},
			},
			want: []Book{handmaidsTale},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)

			if !equalBooks(t, got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestListOtherBooksOnShelf(t *testing.T) {
	type testCase struct {
		books []Book
		index int
		want  []Book
	}

	testCases := map[string]testCase{
		"basic case": {
			books: []Book{handmaidsTale, oryxAndCrake, theBellJar},
			index: 1,
			want:  []Book{handmaidsTale, theBellJar},
		},
		"only one book": {
			books: []Book{handmaidsTale},
			index: 0,
			want:  []Book{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := listOtherBooksOnShelf(tc.index, tc.books)
			if got != tc.want {

			}
		})
	}
}

func Example_main() {
	main()
	// Output
	// Here are the books in common:
	// - The Handmaid's Tale by Margaret Atwood
}

func Example_displayBooks() {
	books := []Book{oryxAndCrake, handmaidsTale}
	displayBooks(books)
	// Output:
	// - Oryx and Crake by Margaret Atwood
	// - The Handmaid's Tale by Margaret Atwood
}

func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i].Author != target[i].Author {
			return false
		}

		if books[i].Title != target[i].Title {
			return false
		}
	}

	return true
}

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range got {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}
