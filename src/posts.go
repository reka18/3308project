package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func postsGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	ok, username := CompareTokens(w, r)

	if !ok {
		return
	}

	RefreshCookie(w, username) /* This updates cookie to restart clock. */
	t := template.Must(template.ParseFiles("web/make_post.html"))
	_ = t.Execute(w, "")
	pushAllResources(w)

}

func postsPOST(w http.ResponseWriter, r *http.Request) {
	CookieDebugger(r, "POST")

	ok, username := CompareTokens(w, r)

	if !ok {
		return
	}

	RefreshCookie(w, username)

	var (
		postTitle = strings.Join(r.Form["post_title"], "")
		postContent = strings.Join(r.Form["post_content"], "")
	)

	newPost := Post{
		Username: username,
		Title:    postTitle,
		Content:  postContent,
	}

	db, _ := Database(DBNAME)
	defer db.Close()

	makePost(newPost, db)
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

func makePost(post Post, db *sql.DB) {

	userid := GetUserId(post.Username, db)

	_, e := db.Exec("INSERT INTO posts (userid, title, content, upvotes, downvotes, deleted) " +
		"VALUES ($1, $2, $3, 0, 0, false);", userid, post.Title, post.Content)
	if e != nil {
		log.Println(Warn("Unable to execute user query."))
		log.Println(Warn(e))
	}

}