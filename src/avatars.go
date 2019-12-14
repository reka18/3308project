package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func avatarGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "AVATAR")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	db, _ := Database(DBNAME)
	defer db.Close()

	bytes := GetAvatar(username, db)

	w.Header().Set("Content-Type", "image/png")
	_, e := w.Write(bytes)
	if e != nil {
		log.Println(Warn("Error writing to response."))
	}

}

func avatarPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "AVATAR")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	_ = r.ParseMultipartForm(10 << 20)

	file, handler, e := r.FormFile("Avatar")
	if e != nil {
		log.Println(Warn("Error retrieving image."), e)
		return
	}
	defer file.Close()

	log.Println(Info("Uploaded file: ", handler.Filename))
	log.Println(Info("File size: ", handler.Size))
	log.Println(Info("MIME Header: ", handler.Header))

	fileBytes, e := ioutil.ReadAll(file)
	if e != nil {
		log.Println("Error reading image bytes.")
	}

	// log.Println("ioutil.ReadAll: ", fileBytes[:10])

	db, _ := Database(DBNAME)

	UpdateAvatar(username, fileBytes, db)
	userPage := fmt.Sprintf("/%s", username)
	http.Redirect(w, r, userPage, http.StatusSeeOther)
}

func AvatarHandler(w http.ResponseWriter, r *http.Request) {

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
		log.Println(Warn("Error retrieving image from database."))
		log.Println(Warn(e))
	}

	log.Println(Info("Database image peek: ", avatarBytes[:10]))

	return avatarBytes

}

func UpdateAvatar(username string, avatar []byte, db *sql.DB) {

	_, e := db.Exec("UPDATE avatars SET avatar=$1 WHERE userid=(SELECT id FROM users WHERE username=$2);",
		avatar, username)
	if e != nil {
		log.Println(Warn("Unable to execute avatar query."))
		log.Println(Warn(e))
	}

}

func NewUserAvatar(username string, db *sql.DB) {

	fileBytes, e := ioutil.ReadFile("web/images/default_avatar.png")
	if e != nil {
		log.Println(Warn("Unable to read default avatar."))
	}

	_, e = db.Exec("INSERT INTO avatars(userid, avatar) VALUES ((SELECT id FROM users WHERE username=$1), $2)", username, fileBytes)
	if e != nil {
		log.Println(Warn("Unable to post default avatar on user creation."))
	}

}