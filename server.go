package main

import (
  "fmt"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Go World")
}



func main() {


  var port string
  var ok  bool
  port, ok = os.LookupEnv("SERVER_PORT")
  // port = os.Getenv("SERVER_PORT")

  if !ok {
      fmt.Println("SERVER_PORT is not set")
      port = "8081"
  }
  fmt.Println("SERVER_PORT is " + port)
  // port = ":8082"
  http.HandleFunc("/", handler) 
  // http.ListenAndServe(":8080", nil)
  port = ":" + port
  http.ListenAndServe(port, nil)
}

