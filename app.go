package main

import (                                   #importing libraries
  "log"
  "net/http"
)

func main() {                                      #FileServer function: responds to all http requests
  fs := http.FileServer(http.Dir("./static"))       #handle function: registers FileServer as handler of all requests
  http.Handle("/", fs)                              #"/" matches all request paths

  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)               #Listen on a port (in this case port 8080) and thus accept connections to internet
}
