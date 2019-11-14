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

	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if e != nil {
		log.Printf(Warn("Unable to hash: %s"), e)
	}
	return string(hash)
}

func VerifyKey(dbPasswordHash string, password string) bool {

	e := bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password))
	if e != nil {
		log.Println(e)
		return false
	}
	return true

}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func AddCookie(w http.ResponseWriter, username string) {

	RefreshCookie(w, username)

	secret := GenerateKey(string(securecookie.GenerateRandomKey(64)))

	cookie := http.Cookie {
		Name:    	"socialmediasite",
		Value:   	fmt.Sprintf("%s:::%s", username, secret),
		MaxAge:		300,
		Expires: 	time.Now().Local().Add(time.Hour * 6),
	}
	http.SetCookie(w, &cookie)

	_, e := redisConn.Do("SET", username, fmt.Sprintf(secret))
	if e != nil {
		log.Printf(Warn("Failed to set cookie for %s."), username)
		log.Println(e)
	}
	_, e = redisConn.Do("EXPIRE", username, http.StatusMultipleChoices)
	if e != nil {
		log.Printf(Warn("Failed to set cookie expiration for %s."), username)
		log.Println(e)
	}

}

func DeleteCookie(w http.ResponseWriter, username string) {

	cookie := http.Cookie {
		Name:		"socialmediasite",
		Value:		"",
		MaxAge:		-1,
		/* Some browsers dont understand `MaxAge` so we add this. */
		Expires: 	time.Now().Add(-100 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	_, e := redisConn.Do("DEL", username)
	if e != nil {
		log.Printf(Warn("Failed to delete cookie on redis server for %s."), username)
	}
	log.Printf(Info("Deleted cookies for '%s'."), username)

}

func RefreshCookie(w http.ResponseWriter, username string) {
	/*
	THIS REFRESHES THE EXPIRATION OF THE COOKIE ON REDIS
	 */
	_, e := redisConn.Do("EXPIRE", username, http.StatusMultipleChoices)
	if e != nil {
		log.Printf(Warn("Failed to refresh cookie expiration for %s."), username)
		log.Println(e)
	}
}

func CompareTokens(w http.ResponseWriter, r *http.Request) (bool, string) {

	log.Printf(Info("Cookie authentication request from: %s"), GetIP(r))

	cookie, _ := r.Cookie("socialmediasite")

	if cookie == nil || cookie.Value == "" {
		log.Println(Warn("Unauthorized access attempt."))
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
			log.Printf(Warn("Unauthorized access attempt with stale cookie for %s."), username)
			DeleteCookie(w, username)
			http.Redirect(w, r, "/logout", http.StatusForbidden)
			return false, ""
		}
		log.Printf(Success("Cookie authentication successful for '%s'."), username)
		return true, username
	}

}

func CookieDebugger(r *http.Request, pagename string) {

	log.Printf(Info("'%s' page cookies: '%v'"), pagename, r.Cookies())

}