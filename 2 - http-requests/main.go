package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	request, err := http.Get("https://www.google.com")
	handleError(err)
	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	handleError(err)

	fmt.Println(string(response))
}
