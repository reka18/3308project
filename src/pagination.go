package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ParseLimit(r *http.Request, defaultLimit int) int {

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
