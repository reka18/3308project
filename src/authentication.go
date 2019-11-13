package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

func GenerateKey(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Unable to hash: %s", err)
	}
	return string(hash)
}

func VerifyKey(dbPasswordHash string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}

func AddCookie(w http.ResponseWriter, username string) {

	secret := GenerateKey(string(securecookie.GenerateRandomKey(64)))

	cookie := http.Cookie {
		Name:    "socialmediasite",
		Value:   fmt.Sprintf("%s:::%s", username, secret),
		Expires: time.Now().Local().Add(time.Hour * 6),
	}
	http.SetCookie(w, &cookie)

	_, _ = redisConn.Do("SET", username, fmt.Sprintf(secret))
	_, _ = redisConn.Do("EXPIRE", username, 60) /* Set to 1 minute for debugging */

}

func DeleteCookie(w http.ResponseWriter, username string) {

	cookie := http.Cookie {
		Name:		"socialmediasite",
		Value:		"",
		MaxAge:		-1,
	}
	http.SetCookie(w, &cookie)

	_, e := redisConn.Do("DEL", username)
	if e != nil {
		log.Println("Failed to delete cookie to redis server.")
	}
	log.Printf("Deleted cookies for '%s'.", username)
}

func CompareTokens(w http.ResponseWriter, r *http.Request) (bool, string) {

	cookie, _ := r.Cookie("socialmediasite")

	if cookie.Value == "" {
		log.Println("Unauthorized access.")
		http.Redirect(w, r, "/logout", http.StatusForbidden)
		return false, ""
	} else {
		values := strings.Split(cookie.Value, ":::")

		username := values[0]
		cookieSecret := values[1]

		/* result is an interface so we can't caste it */
		result, _ := redisConn.Do("GET", username)
		redisSecret := fmt.Sprintf("%s", result)

		if redisSecret != cookieSecret {
			log.Println("Unauthorized access.")
			DeleteCookie(w, username)
			http.Redirect(w, r, "/logout", http.StatusForbidden)
			return false, ""
		}
		log.Println("Cookie authentication successful.")
		return true, username
	}

}