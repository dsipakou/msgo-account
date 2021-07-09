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
	a.Router.HandleFunc("/api/accounts", a.GetAccountsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/users", a.GetUsersHandler()).Methods("GET")
	a.Router.HandleFunc("/api/transactions", a.CreateTransactionHandler()).Methods("POST")
  a.Router.HandleFunc("/api/accounts", a.CreateAccountHandler()).Methods("POST")
  a.Router.HandleFunc("/api/users", a.CreateUserHandler()).Methods("POST")
  a.Router.HandleFunc("/api/transactions", a.DeleteTransactionHandler()).Methods("DELETE")
  a.Router.HandleFunc("/api/accounts", a.DeleteAccountHandler()).Methods("DELETE")
  a.Router.HandleFunc("/api/transactions", a.UpdateTransactionHandler()).Methods("PATCH")
  a.Router.HandleFunc("/api/accounts", a.UpdateAccountHandler()).Methods("PATCH")
}
