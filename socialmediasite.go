package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"html/template"
)

type User struct {
	Name string
}

func usrLandingGET(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("templates/user_landing.html"))
	user := User{
		Name: "Tom Anderson",
	}
	tmpl.Execute(w, user)
}

func usrLandingPOST(w http.ResponseWriter, r *http.Request){
	log.Println("New Post:")
	if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    newpost := r.FormValue("newpost")
    
    log.Println(newpost)
}

func userLandingHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method{

	case "GET":
		usrLandingGET(w,r)
	
	case "POST":
		usrLandingPOST(w,r)
		usrLandingGET(w,r)
	}
}

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

	http.HandleFunc("/user_landing/", userLandingHandler)

	log.Println("Listening...")
	e := http.ListenAndServe(":3000", nil)
	if e != nil {
		log.Fatal(e)
	}
}

