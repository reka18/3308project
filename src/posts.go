package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func postsGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	username := CompareTokens(w, r)

	RefreshCookie(w, username) /* This updates cookie to restart clock. */
	t := template.Must(template.ParseFiles("web/make_post.html"))
	_ = t.Execute(w, "")
	pushAllResources(w)

}

func postsPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	log.Println(Green("At the right page."))

	username := CompareTokens(w, r)

	RefreshCookie(w, username)

	var postContent = r.FormValue("content")


	db, _ := Database(DBNAME)
	defer db.Close()

	makePost(username, postContent, db)

	userPage := fmt.Sprintf("/%s", username)
	http.Redirect(w, r, userPage, http.StatusSeeOther)
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		postsGET(w, r)
	case "POST":
		postsPOST(w, r)
	}
}

func makePost(username string, post string, db *sql.DB) {

	userid := GetUserId(username, db)

	_, e := db.Exec("INSERT INTO posts (userid, content, upvotes, downvotes, deleted) " +
		"VALUES ($1, $2, 0, 0, false);", userid, post)
	if e != nil {
		log.Println(Warn("Unable to execute user query."))
		log.Println(Warn(e))
	}

}