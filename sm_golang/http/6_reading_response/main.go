package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)

	if err != nil{
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println("body byte", bodyBytes)

	bodyText := string(bodyBytes)
	fmt.Println("body text", bodyText)

	max := 250
	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println(bodyText[:max])
}