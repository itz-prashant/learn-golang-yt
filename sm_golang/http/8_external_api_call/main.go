package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type catFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func responseWriter(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (catFactResponse, error) {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)

	if err != nil {
		return catFactResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return catFactResponse{}, fmt.Errorf("external api failed: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return catFactResponse{}, err
	}

	var data catFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return catFactResponse{}, err
	}

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		responseWriter(w, http.StatusMethodNotAllowed, map[string]any{
			"ok" :false,
			"error": "Only Get method is allowed",
		})
	}

	data, err := fetchCatFact()

	if err != nil {
		responseWriter(w, http.StatusBadGateway,map[string]any{
			"ok" :false,
			"error": "Failed to fetch data",
		})
	}

	responseWriter(w, http.StatusOK, map[string]any{
			"ok" :true,
			"timeStamp": time.Now().UTC(),
			"external": map[string]any{
				"fact": data.Fact,
				"length": data.Length,
			},
		})
}

func main() {

	http.HandleFunc("/external", externalHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}
