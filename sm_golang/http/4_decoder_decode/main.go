package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type testRequest struct {
	name string `jaon:"name"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		data := map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		}
		writeJson(w, http.StatusMethodNotAllowed, data)
		return
	}

	defer r.Body.Close()

	var req testRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid json format",
		})
		return
	}

	req.name = strings.TrimSpace(req.name)

	if req.name == "" {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Name must not be empty",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok":   true,
		"data": req,
	})
}

func main() {

	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
