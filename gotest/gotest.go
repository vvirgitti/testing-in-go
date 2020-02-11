package gotest

import "errors"

type Book struct {
	title string
	category string
	pages int
}

type User struct {
	name string
	hasAlreadyBook bool
}

type Library interface {
	BorrowBook(user User) (bool, error)
}

func BorrowBook(user User) (bool, error) {
	if user.hasAlreadyBook == false {
		return true, nil
	} else {
		return false, errors.New("user has already borrowed a book")
	}
}