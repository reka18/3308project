package main

import "net/http"

func BaseUrlHandler(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "BASE")

	http.Redirect(w, r, "/login", 303)
}