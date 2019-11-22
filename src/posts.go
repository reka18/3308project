package main

import (
	"html/template"
	"net/http"
)

func PostsGet(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	ok, username := CompareTokens(w, r)

	if ok {
		RefreshCookie(w, username) /* This updates cookie to restart clock. */
		t := template.Must(template.ParseFiles("web/make_post.html"))
		_ = t.Execute(w, "")
		pushAllResources(w,)
	}

}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	PostsGet(w, r)
}