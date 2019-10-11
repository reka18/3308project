package main

import (
	"fmt"
	"log"
	"net/http"
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
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/user_landing/", userLandingHandler)
	http.Handle("/", fs)
	

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}