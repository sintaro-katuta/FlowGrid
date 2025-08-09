package main

import (
	"fmt"
	"net/http"
	"github.com/sintaro/FlowGrid/backend/internal/config"
	"log"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	// データベースに接続する
	db, err := models.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/health", healthHandler)

	port := "8080"
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
	defer db.Close()
}
