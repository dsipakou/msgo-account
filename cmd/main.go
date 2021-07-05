package main

import (
	"log"
	"msgo-account/pkg/db/models"
  "msgo-account/pkg/db"
	"msgo-account/pkg/repository"
	"net/http"
	"os"
  "fmt"
	"github.com/joho/godotenv"
)

// import "msgo-account/api/server"

func main() {
	// router := server.Init()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot open dotenv file: %s", err.Error())
	}

	rep := repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	log.Printf(rep.Host)
	dbConnect, err := repository.InitDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	dbConnect.Ping()
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Index page")
	})

	t := &models.Transaction{
		UserId:   1,
		Category: "temp_category",
		Amount:   50,
	}

  fmt.Println(t)

  err = db.CreateTransaction(t)
	if err != nil {
		log.Println("Cannot insert record")
	}

	err = http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Println(err)
	}
}
