package main

import (
  "encoding/json"
  "log"
  "net/http"
  "time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
}
