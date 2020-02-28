package main

import "testing"

func TestLibrary(t *testing.T) {
	t.Run("can add a book", func(t *testing.T) {
		library := Library{[]Book{
			{
				category:"manga",
				title: "Naruto",
			},
			{
				category: "play",
				title: "Romeo and Juliet",
			},
		},
		}

		newBook := Book{title:"The three musketeers", category:"novel"}

		got := library.AddBook(newBook)

		if len(got) != 3 {
			t.Errorf("Expected 3 books but got %d", len(got))
		}
	})

	t.Run("can remove a book", func(t *testing.T) {
		library := Library{[]Book{
			{
				category:"manga",
				title: "Naruto",
			},
			{
				category: "play",
				title: "Romeo and Juliet",
			},
		},
		}

		got := library.RemoveBook(Book{category:"manga", title:"Naruto"})

		if len(got) != 1 {
			t.Errorf("Expected 1 books but got %d", len(got))
		}
	})
}

func TestUser(t *testing.T) {
	t.Run("can borrow a book", func(t *testing.T) {
		harry := Book{title: "Harry Potter", category: "novel"}
		bob := User{
			name:            "bob",
			hasBorrowedBook: false,
			borrowedBook:    harry,
		}

		library := Library{[]Book{
			{
				title: "20th century boys",
				category: "manga",
			},
			{
				title: "Harry Potter",
				category: "novel",
			},
		}}

		BorrowBook(bob, harry)

		got := len(library.Books)

		if got != 1 {
			t.Errorf("")
		}


	})
}