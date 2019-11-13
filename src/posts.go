package main

import (
	"html/template"
	"log"
	"net/http"
)

func PostsGet(w http.ResponseWriter, r *http.Request) {

	log.Println("Posts page get cookies: ", r.Cookies())

	ok, _ := CompareTokens(w, r)

	if ok {
		t := template.Must(template.ParseFiles("web/make_post.html"))
		_ = t.Execute(w, "")
		pushAllResources(w,)
	}

}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	PostsGet(w, r)
}