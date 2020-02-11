package gotest

import "testing"

func TestBorrowBook(t *testing.T) {
	bob := User{name: "Bob", hasAlreadyBook: false}

	got, err := BorrowBook(bob)
	if err != nil {
		t.Errorf("BorrowBook %d; want true", got)
	}
}