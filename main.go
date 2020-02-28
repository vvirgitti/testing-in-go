package main

import (
	"log"
	"net/http"
)

func main () {
	handler := http.HandlerFunc(Server)

	if err := http.ListenAndServe(":2000", handler); err != nil {
		log.Fatalf("Could not listen on port 2000 %v", err)
	}
}
