package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	// EXAMPLE USAGE
	//AddNewUserAccount(40, "Rodrigo", "Garcia", "rigo.garcia@colorado.edu",
	//	"M", true, "iamtheverymodelofthemodernmajorgeneral")
	//AddNewUserAccount(36, "Reagan", "Karnes", "reagan.karnes@colorado.edu",
	//	"M", true, "abcdefghijklmnopqrstuvwxyz1234567890")

	// EXAMPLE USAGE
	result := LoginUserAccount("reagan.karnes@colorado.edu", "abcdefghijklmnopqrstuvwxyz1234567890")
	PrintUser(result)

	fs := http.FileServer(http.Dir("source"))
	http.Handle("/", fs)

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

