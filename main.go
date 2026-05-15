package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

func handlerAPI(w http.ResponseWriter, r *http.Request) {
	var data Request
	json.NewDecoder(r.Body).Decode(&data)

	json.NewEncoder(w).Encode(map[string]string{"message": data.Message})
}

func main() {

	http.HandleFunc("/api", handlerAPI)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port :8080")
}
