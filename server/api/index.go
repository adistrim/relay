package handler

import (
	"net/http"
	"relay/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := routes.MainRouter()
	router.ServeHTTP(w, r)
}
