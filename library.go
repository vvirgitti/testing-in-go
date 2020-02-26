package main

type Book struct {
	title string
	category string
}

type Library struct {
	Books []Book
}

type LibraryStore interface {
	AddBook(book Book) []Book
	RemoveBook(book Book) []Book
}

func (l Library) AddBook(book Book) []Book {
	l.Books = append(l.Books, book)
	return l.Books
}

func (l Library) RemoveBook(book Book) []Book {
	var newBooks []Book
	for i, b := range l.Books {
		if b == book {
			l.Books[i] = l.Books[len(l.Books)-1]
		}
		newBooks =l.Books[:len(l.Books)-1]
	}
	return newBooks
}

type User struct {
	name string
	hasBorrowedBook bool
	borrowedBook Book
}









