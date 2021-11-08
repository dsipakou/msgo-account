package server

import (
	"msgo-account/pkg/db"

	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
	DB     db.GeneralDB
}

func Init() *Api {
	a := &Api{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *Api) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/transactions", a.GetTransactionsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/transactions", a.GetTransactionsHandler()).Methods("GET").Queries("sorting", "{[a-zA-Z]+}", "dateFrom", "{[0-9-]+}", "dateTo", "{[0-9-]+}")
	a.Router.HandleFunc("/api/transactions/month/{dateFrom:[0-9-]+}/{dateTo:[0-9-]+}", a.GetGroupedTransactionsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/transactions", a.CreateTransactionHandler()).Methods("POST")
	a.Router.HandleFunc("/api/transactions", a.DeleteTransactionHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/transactions", a.UpdateTransactionHandler()).Methods("PATCH")
	a.Router.HandleFunc("/api/accounts", a.GetAccountsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/accounts", a.CreateAccountHandler()).Methods("POST")
	a.Router.HandleFunc("/api/accounts", a.DeleteAccountHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/accounts", a.UpdateAccountHandler()).Methods("PATCH")
	a.Router.HandleFunc("/api/login", a.UserLoginHandler()).Methods("POST")
	a.Router.HandleFunc("/api/users", a.GetUsersHandler()).Methods("GET")
	a.Router.HandleFunc("/api/users", a.CreateUserHandler()).Methods("POST")
	a.Router.HandleFunc("/api/users", a.DeleteUserHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/users", a.UpdateUserHandler()).Methods("PATCH")
	a.Router.HandleFunc("/api/categories", a.GetCategoriesHandler()).Methods("GET")
	a.Router.HandleFunc("/api/categories", a.CreateCategoryHandler()).Methods("POST")
	a.Router.HandleFunc("/api/categories", a.DeleteCategoryHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/categories", a.UpdateCategoryHandler()).Methods("PATCH")
	a.Router.HandleFunc("/api/currencies", a.GetCurrenciesHandler()).Methods("GET")
	a.Router.HandleFunc("/api/currencies", a.CreateCurrencyHandler()).Methods("POST")
	a.Router.HandleFunc("/api/currencies", a.DeleteCurrencyHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/currencies", a.UpdateCurrencyHandler()).Methods("PATCH")
	a.Router.HandleFunc("/api/rates", a.GetRatesHandler()).Methods("GET")
	a.Router.HandleFunc("/api/rates", a.CreateRateHandler()).Methods("POST")
	a.Router.HandleFunc("/api/rates", a.DeleteRateHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/rates", a.UpdateRateHandler()).Methods("PATCH")
	a.Router.HandleFunc("/api/budget", a.GetBudgetHandler()).Methods("GET")
	a.Router.HandleFunc("/api/budget", a.GetBudgetForPeriodHandler()).Methods("GET").Queries("dateFrom", "{[0-9-]+}", "dateTo", "{[0-9-]+}")
	a.Router.HandleFunc("/api/budget/period", a.GetBudgetUsageForPeriodHandler()).Methods("GET")
	a.Router.HandleFunc("/api/budget", a.CreateBudgetHandler()).Methods("POST")
	a.Router.HandleFunc("/api/budget", a.DeleteBudgetHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/budget", a.UpdateBudgetHandler()).Methods("PATCH")
}
