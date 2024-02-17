package main

import (
  "fmt"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Go World 2")
}



func main() {

  var port string
  var ok  bool
  port, ok = os.LookupEnv("SERVER_PORT")
  // port = os.Getenv("SERVER_PORT")

  if !ok {
      fmt.Println("SERVER_PORT is not set")
      port = "8080" // default
  }

  fmt.Println("SERVER_PORT is " + port)

  http.HandleFunc("/", handler)
  port = ":" + port
  http.ListenAndServe(port, nil)
}

