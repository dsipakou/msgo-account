package main

import "net/http"
import "log"

func main() {
  http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
    log.Println("Index page")
  })

  http.ListenAndServe(":9091", nil)
}
