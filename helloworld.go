package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handlePostRequest)
	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData JsonRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		httpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if requestData.Message == "" {
		httpError(w, "Missing or altered 'message' field", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message: %s\n", requestData.Message)

	response := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func httpError(w http.ResponseWriter, message string, code int) {
	errorResponse := JsonResponse{
		Status:  fmt.Sprintf("%d", code),
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse)
}
