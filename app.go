package main

import (                                   #importing libraries
  "log"
  "net/http"
)

func main() {                                      
  fs := http.FileServer(http.Dir("./static"))       #FileServer function: responds to all http requests
  http.Handle("/", fs)                              #handle function: registers FileServer as handler of all requests
                                                                                      #"/" matches all request paths
  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)               #Listen on a port (in this case port 8080) and thus accept connections to internet
}
