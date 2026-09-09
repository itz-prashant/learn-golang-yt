package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type catFactResponse struct {
	Fact   string `json:"fact"`
	Length int `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Read body fail", err)
		return
	}

	var data catFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("Json unmarshel fail", err)
		return
	}

	fmt.Println(data.Fact, data.Length)
}
