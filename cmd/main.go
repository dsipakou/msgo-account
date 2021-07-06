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

	t := &models.TransactionRequest{
		UserId:   1,
		Category: "temp_category",
		Amount:   50,
	}

	fmt.Println(t)

  transactions, err := api.DB.GetTransactions()
	check(err)

  http.HandleFunc("/", api.Router.ServeHTTP)
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
