package main

import (
	"reflect"
	"testing"
)

func TestRecommendOtherBooks(t *testing.T) {
	type testCase struct {
		bookworms []Bookworm
		want      []Bookworm
	}

	tests := map[string]testCase{
		"empty bookworms": {
			bookworms: []Bookworm{},
			want:      []Bookworm{},
		},
		"single bookworm with no books": {
			bookworms: []Bookworm{{Name: "Alice", Books: []Book{}}},
			want:      []Bookworm{{Name: "Alice", Books: []Book{}}},
		},
		"single bookworm with one book": {
			bookworms: []Bookworm{{Name: "Alice", Books: []Book{theBellJar}}},
			want:      []Bookworm{{Name: "Alice", Books: []Book{}}},
		},
		"single bookworm with multiple books": {
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake}},
			},
			want: []Bookworm{
				{Name: "Alice", Books: []Book{}},
			},
		},
		"two bookworms with overlapping books": {
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake}},
				{Name: "Bob", Books: []Book{theBellJar, handmaidsTale}},
			},
			want: []Bookworm{
				{Name: "Alice", Books: []Book{handmaidsTale}},
				{Name: "Bob", Books: []Book{oryxAndCrake}},
			},
		},
		"two bookworms with no overlapping books": {
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar}},
				{Name: "Bob", Books: []Book{oryxAndCrake}},
			},
			want: []Bookworm{
				{Name: "Alice", Books: []Book{}},
				{Name: "Bob", Books: []Book{}},
			},
		},
		"two bookworms with one overlapping book one has nothing else on their shelf": {
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake}},
				{Name: "Bob", Books: []Book{theBellJar}},
			},
			want: []Bookworm{
				{Name: "Alice", Books: []Book{}},
				{Name: "Bob", Books: []Book{oryxAndCrake}},
			},
		},
		"three bookworms with multiple books in common": {
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake, handmaidsTale}},
				{Name: "Bob", Books: []Book{theBellJar, handmaidsTale, janeEyre}},
				{Name: "Charlie", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: []Bookworm{
				{Name: "Alice", Books: []Book{janeEyre}},
				{Name: "Bob", Books: []Book{oryxAndCrake}},
				{Name: "Charlie", Books: []Book{theBellJar}},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := recommendOtherBooks(tc.bookworms)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got different bookworms than expected: %v, want %v", got, tc.want)
			}
		})
	}
}

func TestBuildBookRecommendations(t *testing.T) {
	tests := []struct {
		name      string
		bookworms []Bookworm
		want      bookRecommendations
	}{
		{
			name:      "empty bookworms",
			bookworms: []Bookworm{},
			want:      bookRecommendations{},
		},
		{
			name: "single bookworm with no books",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{}},
			},
			want: bookRecommendations{},
		},
		{
			name: "single bookworm with one book",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar}},
			},
			want: bookRecommendations{},
		},
		{
			name: "single bookworm with multiple books",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake}},
			},
			want: bookRecommendations{
				theBellJar: bookCollection{
					oryxAndCrake: struct{}{},
				},
				oryxAndCrake: bookCollection{
					theBellJar: struct{}{},
				},
			},
		},
		{
			name: "two bookworms with overlapping books",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake}},
				{Name: "Bob", Books: []Book{theBellJar, handmaidsTale}},
			},
			want: bookRecommendations{
				theBellJar: bookCollection{
					oryxAndCrake:  struct{}{},
					handmaidsTale: struct{}{},
				},
				oryxAndCrake: bookCollection{
					theBellJar: struct{}{},
				},
				handmaidsTale: bookCollection{
					theBellJar: struct{}{},
				},
			},
		},
		{
			name: "two bookworms with no overlapping books",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar}},
				{Name: "Bob", Books: []Book{oryxAndCrake}},
			},
			want: bookRecommendations{},
		},
		{
			name: "three bookworms with complex relationships",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake, handmaidsTale}},
				{Name: "Bob", Books: []Book{theBellJar, handmaidsTale, janeEyre}},
				{Name: "Charlie", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: bookRecommendations{
				theBellJar: bookCollection{
					oryxAndCrake:  struct{}{},
					handmaidsTale: struct{}{},
					janeEyre:      struct{}{},
				},
				oryxAndCrake: bookCollection{
					theBellJar:    struct{}{},
					handmaidsTale: struct{}{},
					janeEyre:      struct{}{},
				},
				handmaidsTale: bookCollection{
					theBellJar:   struct{}{},
					oryxAndCrake: struct{}{},
					janeEyre:     struct{}{},
				},
				janeEyre: bookCollection{
					theBellJar:    struct{}{},
					oryxAndCrake:  struct{}{},
					handmaidsTale: struct{}{},
				},
			},
		},
		{
			name: "bookworm with duplicate books",
			bookworms: []Bookworm{
				{Name: "Alice", Books: []Book{theBellJar, oryxAndCrake, theBellJar}},
			},
			want: bookRecommendations{
				theBellJar: bookCollection{
					oryxAndCrake: struct{}{},
					theBellJar:   struct{}{},
				},
				oryxAndCrake: bookCollection{
					theBellJar: struct{}{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildBookRecommendations(tt.bookworms)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildBookRecommendations() = %v, want %v", got, tt.want)
			}
		})
	}
}
