package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := returnError(true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	
}

type specialError struct {
	errorMessage string
	errorCode int
}

func (s specialError) Error() string {
	return s.errorMessage + " " +strconv.Itoa(s.errorCode)
}

func returnError(returnError bool) (string, error) {
	if returnError {
		return "", specialError{errorMessage: "Content Not Found!", errorCode: 404}
	}
	return "No error", nil
}
