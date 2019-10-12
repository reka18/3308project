package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("source"))
	http.Handle("/", fs)

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}