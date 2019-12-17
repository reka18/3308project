package main

import (
	"net/http"
)

func BaseUrlHandler(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "BASE PAGE REDIRECT")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}