package main

import (
	"database/sql"
	"log"
	"net/http"
)

func avatarGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "AVATAR")

	username := CompareTokens(w, r)
	RefreshCookie(w, username) /* This updates cookie to restart clock. */

	db, _ := Database(DBNAME)
	defer db.Close()
	code, _ := w.Write(GetAvatar(username, db))
	log.Println(Info("Avatar Response: ", code))

}

func avatarPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "AVATAR")

	username := CompareTokens(w, r)
	RefreshCookie(w, username) /* This updates cookie to restart clock. */

	avatarBytes := r.FormValue("new_avatar")

	db, _ := Database(DBNAME)
	defer db.Close()

	PostAvatar(username, []byte(avatarBytes), db)

}

func AvatarHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		avatarGET(w, r)
	case "POST":
		avatarPOST(w, r)
	}

}

func GetAvatar(username string, db *sql.DB) []byte {

	var (
		avatarid	int
		userid		int
		avatarBytes	[]byte
	)

	r := db.QueryRow("SELECT * FROM avatars WHERE userid=(SELECT id FROM users WHERE username=$1);", username)

	e := r.Scan(&avatarid, &userid, &avatarBytes)

	if e != nil {
		return nil
	}

	return avatarBytes

}

func PostAvatar(username string, avatar []byte, db *sql.DB) {

	userid := GetUserId(username, db)
	_, e := db.Exec("UPDATE avatars SET avatar=$1 WHERE userid=$2;", avatar, userid)
	if e != nil {
		log.Println(Warn("Unable to execute avatar query."))
		log.Println(Warn(e))
	}

}