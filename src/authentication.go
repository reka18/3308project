package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Unable to hash: %s", err)
	}
	return string(hash)
}

func VerifyPW(dbPasswordHash string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}

func SecureCookieController(w http.ResponseWriter, username string) {
	key := securecookie.GenerateRandomKey(64)
	c := http.Cookie {
		Name:	fmt.Sprintf("socialmediasite:login-cookie:[%s]", username),
		Value:	string(key),
	}
	log.Printf("Cookie: %v", &c)
	http.SetCookie(w, &c)
}