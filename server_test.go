package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLibrary_Server(t *testing.T) {
	t.Run("GET / returns Library rules", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Server(response, request)

		got := response.Body.String()
		want := "Library rules"

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}