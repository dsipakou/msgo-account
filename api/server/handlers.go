package server

import (
	"fmt"
	"net/http"
)

func (a *Api) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Account API")
	}
}
