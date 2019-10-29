package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	// RESETS THE DATABASE TO AN EMPTY STATE
	if len(os.Args) > 1 {
		db, _ := Database(PGNAME)

		if os.Args[1] == "--reset" {
			log.Println("Manually dropping tables.")
			e := DropTables(db)
			if e != nil {
				log.Fatal("Unable to drop tables:", e)
			}
		}

		if os.Args[1] == "--create" {
			e := CreateDatabase(db)
			if e != nil {
				log.Fatal("Unable to create database:", e)
			}
			e = InitializeDatabase(db)
			if e != nil {
				log.Fatal("Unable to initialize database:", e)
			}
		}
	}
	
	fs := http.FileServer(http.Dir("source"))
	http.Handle("/", fs)

	log.Println("Listening...")
	e := http.ListenAndServe(":3000", nil)

	if e != nil {
		log.Fatal(e)
	}
}

