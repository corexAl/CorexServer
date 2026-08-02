package main

import (
	"fmt"
	"net/http"

	"corex-server/api"
)

func main() {
	mux := http.NewServeMux()

	api.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("COREX Server running on :8080")

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
