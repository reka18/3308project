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

func AddCookie(w http.ResponseWriter, username string) {

	secret := GenerateKey(string(securecookie.GenerateRandomKey(64)))

	cookie := http.Cookie {
		Name:    	"screebit",
		Value:   	fmt.Sprintf("%s:%s", username, secret),
		Expires: 	time.Now().Local().Add(time.Hour * 6),
	}
	http.SetCookie(w, &cookie)

	_, e := redisConn.Do("SET", username, fmt.Sprintf(secret))
	if e != nil {
		log.Printf(Fail("Redis failed to set %s cookie."), username)
		log.Println(e)
	}
	_, e = redisConn.Do("EXPIRE", username, 300)
	if e != nil {
		log.Printf(Warn("Redis failed to set '%s' cookie expiration."), username)
		log.Println(e)
	}

}

func DeleteCookie(username string) {

	_, e := redisConn.Do("DEL", username)
	if e != nil {
		log.Printf(Warn("Redis failed to delete %s cookie."), username)
	}
	log.Printf(Info("Redis deleted '%s' cookies."), username)

}

func RefreshCookie(username string) {
	/*
	THIS REFRESHES THE EXPIRATION OF THE COOKIE ON REDIS
	 */
	_, e := redisConn.Do("EXPIRE", username, 300)
	if e != nil {
		log.Printf(Warn("Redis failed to refresh %s cookie for."), username)
		log.Println(e)
		return
	}
	log.Printf(Success("Redis refreshed '%s' cookie.."), username)

}

func CompareTokens(w http.ResponseWriter, r *http.Request) (string, bool) {

	cookie, _ := r.Cookie("screebit")

	if cookie == nil || cookie.Value == "" {
		log.Println(Warn("Unauthorized access attempt."))
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return "", false
	}

	values := strings.Split(cookie.Value, ":")

	username := values[0]
	cookieSecret := values[1]

	/* result is an interface so we can't caste it */
	result, e := redisConn.Do("GET", username)
	if e != nil || fmt.Sprintf("%s", result) != cookieSecret {
		log.Printf(Fail("Redis failed to find `%s` cookie.", username))
		forceLogout(username, w, r)
		return "", false
	}
	log.Printf(Success("Cookie authentication successful for '%s'."), username)
	return username, true

}

func CookieDebugger(r *http.Request, pagename string) {

	log.Printf(Info("'%s' page cookies: '%v'"), pagename, r.Cookies())

}

func forceLogout(username string, w http.ResponseWriter, r *http.Request) {

	log.Printf(Warn("Unauthorized access attempt with stale cookie for '%s'."), username)
	DeleteCookie(username)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}