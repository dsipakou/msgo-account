package main

import "net/http"
import "log"
// import "msgo-account/api/server"

func main() {
  // router := server.Init()

  http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
    log.Println("Index page")
  })

  err := http.ListenAndServe(":9091", nil)

  if (err != nil) {
    log.Println(err)
  }
}
