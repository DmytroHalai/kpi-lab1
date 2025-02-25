package main

import (
	"fmt"
	"net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/time", TimeHandler)

  server := &http.Server{
    Addr:    ":8795",
    Handler: mux,
  }

  server.ListenAndServe()
  fmt.Println("Server is running at http://localhost:8795/time")
}