package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"msgo-account/api/server"
	"msgo-account/pkg/db"
	"net/http"
	"os"
)

func main() {
	api := server.Init()
	api.DB = &db.DB{}
	err := api.DB.Open()
	check(err)

	defer api.DB.Close()

	fmt.Println(api.Router)
	headers := handlers.AllowedHeaders([]string{"Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// http.HandleFunc("/", api.Router.ServeHTTP)
	err = http.ListenAndServe(":9091", handlers.CORS(headers, methods, origins)(api.Router))
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
