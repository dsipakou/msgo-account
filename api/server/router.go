package server

import "github.com/gorilla/mux"
import "fmt"

type Api struct {
  Router *mux.Router
}

func Init() *Api {
  fmt.Println("Hello from Init")
  return &Api{
    Router: mux.NewRouter(),
  }
}
