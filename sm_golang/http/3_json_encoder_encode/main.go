package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok" : true,
		"message" : "json encoded successfully",
		"time" : time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}

func main() {

	http.HandleFunc("/", successHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
