package main

import "net/http"

func BaseUrlHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 303)
}