package main

import (
	"log"
	"net/http"
)

func BaseUrlHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf(Info("Request to BaseUrlHandler from: %s"), GetIP(r))

	CookieDebugger(r, "BASE")

	http.Redirect(w, r, "/login", 303)
}