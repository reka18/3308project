package main

import (
	"log"
	"net/http"

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

func CookieReader(c *http.Cookie, e error) {
	if e != nil {
		log.Printf("Unable to read cookie. Aborting login. %v", e)
		return
	}
	log.Println(c.Name)
	log.Println(c.Value)
	log.Println(c.Raw)
	log.Println(c.RawExpires)
}