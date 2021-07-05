package main

import (
	"fmt"
	"log"
	"msgo-account/api/server"
	"msgo-account/pkg/db"
	"msgo-account/pkg/db/models"
	"net/http"
	"os"
)

func main() {
	api := server.Init()
	api.DB = &db.DB{}
	err := api.DB.Open()
	check(err)

  fmt.Println(api.DB)

	defer api.DB.Close()
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Index page")
	})

	t := &models.Transaction{
		UserId:   1,
		Category: "temp_category",
		Amount:   50,
	}

	fmt.Println(t)

  transactions, err := api.DB.GetTransactions()
	check(err)

  fmt.Println(transactions)
	err = http.ListenAndServe(":9091", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
