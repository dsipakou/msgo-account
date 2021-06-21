package router

import 'github.com/gorilla/mux'

type Api struct {
  Router *mux.Router
}

func Init() *Api {
  return &Api{
    Router: mux.NewRouter(),
  }
}
