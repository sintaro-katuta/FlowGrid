package main

import (
	"fmt"
	"net/http"
	"github.com/sintaro/FlowGrid/backend/models"
	"log"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	port := "8080"
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
