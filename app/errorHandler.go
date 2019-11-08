package main

import "log"

func FailError(err error, message string) {
	if err != nil {
		log.Fatalf("%s : %s", message, err)
	}
}

func PassError(err error, message string) {
	if err != nil {
		log.Printf("%s : %s", message, err)
	}
}