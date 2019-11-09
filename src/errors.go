package main

import "fmt"

type EmptyStringError struct {}

func (e *EmptyStringError) Error() string {
	return fmt.Sprint("empty string error")
}