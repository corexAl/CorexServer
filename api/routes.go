package api

import (
	"encoding/json"
	"net/http"
)

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/health", health)

	mux.HandleFunc("/api/chat", chat)
}

func health(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{
		"status": "online",
	})
}

func chat(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"POST required",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var request ChatRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	// Temporary response.
	// This will later call COREX AI.

	response := ChatResponse{
		Response: "COREX received: " + request.Prompt,
	}

	json.NewEncoder(w).Encode(response)
}
