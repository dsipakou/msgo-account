package main

import (
  "net/http"
  "log"
  "github.com/joho/godotenv"
  "os"
  "msgo-account/pkg/repository"
)
// import "msgo-account/api/server"

func main() {
  // router := server.Init()

  if err := godotenv.Load(); err != nil {
    log.Fatalf("Cannot open dotenv file: %s", err.Error())
  }

  db, err := repository.NewPostgresDB(repository.Config{
    Host: os.Getenv("DB_HOST"),
    Port: os.Getenv("DB_PORT"),
    Username: os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PASSWORD"),
    DBName: os.Getenv("DB_NAME"),
    SSLMode: os.Getenv("DB_SSL_MODE"),
  })

  if err != nil {
    log.Fatalf(err.Error())
  }

  db.Ping()
  http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
    log.Println("Index page")
  })

  err1 := http.ListenAndServe(":9091", nil)

  if (err1 != nil) {
    log.Println(err1)
  }
}
