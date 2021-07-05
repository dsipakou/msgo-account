package server

import (
	"fmt"
	"msgo-account/pkg/db"

	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
	DB     db.GeneralDB
}

func Init() *Api {
	fmt.Println("Hello from Init")
	a := &Api{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *Api) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/transactions", a.GetTransactionsHandler()).Methods("GET")
}
