package main

import (
	_ "SocialMediaSite/handlers"
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
			_ = DropTables(db)
		}

		if os.Args[1] == "--create" {
			log.Println("Manually creating database and initializing tables.")
			_ = CreateDatabase(db)
			_ = InitializeDatabase(db)
		}

		defer db.Close()
	}

	fs := http.FileServer(http.Dir("source"))
	http.Handle("/", fs)

	//http.HandleFunc("/user_landing/", user_landing.UserLandingHandler)

	log.Println("Listening...")
	e := http.ListenAndServe(":3000", nil)
	if e != nil {
		log.Fatal(e)
	}
}

