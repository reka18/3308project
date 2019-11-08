package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/user_landing.html"))
	user := User{
		Firstname: "John",
		Lastname: "Smith",
	}
	PassError(tmpl.Execute(w, user), "Error executing.")
}

func usrLandingPOST(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
       fmt.Fprintf(w, "ParseForm() err: %v", err)
       return
   }
   username := r.FormValue("username")

   log.Println(username)
}


func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		usrLandingGET(w,r)

	case "POST":
		usrLandingPOST(w,r)
		usrLandingGET(w,r)
	}
}