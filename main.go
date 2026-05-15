package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func handlerAPI(w http.ResponseWriter, r *http.Request) {
	var data Request
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if data.Message == "" {
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}
	res := Response{Message: fmt.Sprintf("Response: %s", data.Message)}

	json.NewEncoder(w).Encode(res)
}

func main() {

	http.HandleFunc("/api", handlerAPI)
	fmt.Println("Server started on port :8080")
	http.ListenAndServe(":8080", nil)
}
