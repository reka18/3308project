package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ParseLimitQuery(r *http.Request, defaultLimit int) int {

	value, ok := r.URL.Query()["limit"]
	if !ok {
		log.Println(Warn("No limit value passed in request."))
	} else {
		log.Println(Success(fmt.Sprintf("Limit value of %v detected.", value)))
		parsedLimit, e := strconv.Atoi(value[0])

		if e != nil {
			log.Println(Warn("Unable to parse page limit."))
		} else {
			log.Println(Success("Successfully parsed page limit."))
			defaultLimit = parsedLimit // only set this if no errors
		}
	}

	return defaultLimit

}


func ParseVoteQuery(r *http.Request) (string, int) {

	values, ok := r.URL.Query()["cast"]
	if !ok {
		log.Println(Warn("No cast value passed in request."))
	} else {
		postvote := strings.Split(values[0], "-")
		vote := postvote[0]
		postid, e := strconv.Atoi(postvote[1])
		if e != nil {
			return "", -1
		}

		if vote == "up" || vote == "down" {
			return vote, postid
		}
	}
	return "", -1
}

func ParseAvatarQuery(r *http.Request) string {

	values, ok := r.URL.Query()["user"]
	if !ok {
		log.Println(Warn("No avatar query terms specified."))
	} else {
		log.Println(Info("Found avatar name: ", values))
		username := values[0]
		return strings.Split(username, " ")[0]
	}
	return ""
}

func ParseUserQuery(r *http.Request) string {

	values, ok := r.URL.Query()["user"]
	if !ok {
		log.Println(Warn("No user value passed in request."))
	} else {
		username := strings.Split(values[0], " ")[0]
		return username
	}
	return ""

}