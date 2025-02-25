package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler)

	server := &http.Server{
		Addr:    ":8795",
		Handler: mux,
	}

	fmt.Println("Server is running at http://localhost:8795/time")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
